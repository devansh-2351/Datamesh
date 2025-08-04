package data_availability

import (
	"context"
	"testing"
)

func TestPostAndSampleBlob(t *testing.T) {
	ctx := context.Background()
	engine := &ErasureKZGEngine{}
	data := []byte("test-blob")
	id, err := engine.PostBlob(ctx, data)
	if err != nil || id == "" {
		t.Fatalf("PostBlob failed: %v, id: %s", err, id)
	}

	sampled, err := engine.SampleBlob(ctx, id)
	if err != nil || string(sampled) != string(data) {
		t.Fatalf("SampleBlob failed: %v, data: %s", err, string(sampled))
	}
}

func TestSampleBlobNotFound(t *testing.T) {
	ctx := context.Background()
	engine := &ErasureKZGEngine{}
	id := "nonexistent"
	data, err := engine.SampleBlob(ctx, id)
	if err != nil {
		t.Fatalf("Expected nil error for not found, got: %v", err)
	}
	if data != nil {
		t.Fatalf("Expected nil data for not found, got: %v", data)
	}
} 