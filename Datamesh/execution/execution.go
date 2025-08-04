package execution

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
)

// Transaction types
const (
	TxTypeSet = "set"
	TxTypeGet = "get"
)

// KVTx is a simple key-value transaction
// {"type":"set", "key":"foo", "value":"bar"}
// {"type":"get", "key":"foo"}
type KVTx struct {
	Type  string `json:"type"`
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

// KVStore is a thread-safe in-memory key-value store
var kvStore = struct {
	m map[string]string
	sync.RWMutex
}{m: make(map[string]string)}

// VM defines the interface for a pluggable virtual machine.
type VM interface {
	ExecuteTx(ctx context.Context, tx []byte) ([]byte, error)
	Name() string
}

// EVM is a stub for the Ethereum Virtual Machine.
type EVM struct{}

func (e *EVM) ExecuteTx(ctx context.Context, tx []byte) ([]byte, error) {
	// For demo, use the real key-value logic
	return executeKVTx(ctx, tx)
}

func (e *EVM) Name() string {
	return "EVM (KVStore demo)"
}

// WASM is a stub for the WebAssembly VM.
type WASM struct{}

func (w *WASM) ExecuteTx(ctx context.Context, tx []byte) ([]byte, error) {
	// TODO: Implement WASM execution
	return nil, errors.New("WASM not implemented")
}

func (w *WASM) Name() string {
	return "WASM (stub)"
}

// executeKVTx parses and executes a key-value transaction
func executeKVTx(ctx context.Context, tx []byte) ([]byte, error) {
	var t KVTx
	if err := json.Unmarshal(tx, &t); err != nil {
		return nil, errors.New("invalid tx format")
	}
	switch t.Type {
	case TxTypeSet:
		kvStore.Lock()
		kvStore.m[t.Key] = t.Value
		kvStore.Unlock()
		return []byte("OK"), nil
	case TxTypeGet:
		kvStore.RLock()
		val, ok := kvStore.m[t.Key]
		kvStore.RUnlock()
		if !ok {
			return nil, errors.New("key not found")
		}
		return []byte(val), nil
	default:
		return nil, errors.New("unknown tx type")
	}
} 