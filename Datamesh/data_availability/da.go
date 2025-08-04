package data_availability

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"sync"
)

// DAEngine defines the interface for data availability operations.
type DAEngine interface {
	PostBlob(ctx context.Context, data []byte) (string, error)
	SampleBlob(ctx context.Context, id string) ([]byte, error)
}

// In-memory blob storage
var blobStore = struct {
	m map[string][]byte
	sync.RWMutex
}{m: make(map[string][]byte)}

// ErasureKZGEngine is a stub for erasure coding and KZG commitments.
type ErasureKZGEngine struct{}

func (e *ErasureKZGEngine) PostBlob(ctx context.Context, data []byte) (string, error) {
	// For demo, store the blob and return its hash as ID
	h := sha256.Sum256(data)
	id := hex.EncodeToString(h[:])
	blobStore.Lock()
	blobStore.m[id] = data
	blobStore.Unlock()
	return id, nil
}

func (e *ErasureKZGEngine) SampleBlob(ctx context.Context, id string) ([]byte, error) {
	blobStore.RLock()
	data, ok := blobStore.m[id]
	blobStore.RUnlock()
	if !ok {
		return nil, nil // Not found
	}
	return data, nil
} 