package consensus

import (
	"context"
	"fmt"
	"sync"
	datameshpb "github.com/datamesh-labs/datamesh/proto"
)

// ConsensusEngine defines the interface for consensus operations.
type ConsensusEngine interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Status() string
	IncrementBlock()
	AddValidator(v *datameshpb.Validator) bool
	RemoveValidator(id string) bool
	ListValidators() []*datameshpb.Validator
}

// TendermintEngine is a stub for Tendermint/CometBFT integration.
type TendermintEngine struct {
	mu         sync.RWMutex
	height     int64
	running    bool
	validators map[string]*datameshpb.Validator
}

func (t *TendermintEngine) Start(ctx context.Context) error {
	t.mu.Lock()
	t.running = true
	t.mu.Unlock()
	return nil
}

func (t *TendermintEngine) Stop(ctx context.Context) error {
	t.mu.Lock()
	t.running = false
	t.mu.Unlock()
	return nil
}

func (t *TendermintEngine) Status() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return fmt.Sprintf("Tendermint consensus engine (stub), running: %v, block height: %d, validators: %d", t.running, t.height, len(t.validators))
}

func (t *TendermintEngine) IncrementBlock() {
	t.mu.Lock()
	t.height++
	t.mu.Unlock()
}

func (t *TendermintEngine) AddValidator(v *datameshpb.Validator) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.validators == nil {
		t.validators = make(map[string]*datameshpb.Validator)
	}
	if _, exists := t.validators[v.Id]; exists {
		return false
	}
	t.validators[v.Id] = v
	return true
}

func (t *TendermintEngine) RemoveValidator(id string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.validators == nil {
		return false
	}
	if _, exists := t.validators[id]; !exists {
		return false
	}
	delete(t.validators, id)
	return true
}

func (t *TendermintEngine) ListValidators() []*datameshpb.Validator {
	t.mu.RLock()
	defer t.mu.RUnlock()
	vals := make([]*datameshpb.Validator, 0, len(t.validators))
	for _, v := range t.validators {
		vals = append(vals, v)
	}
	return vals
} 