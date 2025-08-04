package execution

import (
	"context"
	"encoding/json"
	"testing"
)

func TestKVStoreSetAndGet(t *testing.T) {
	ctx := context.Background()
	vm := &EVM{}

	// Set a value
	txSet := KVTx{Type: TxTypeSet, Key: "foo", Value: "bar"}
	b, _ := json.Marshal(txSet)
	res, err := vm.ExecuteTx(ctx, b)
	if err != nil || string(res) != "OK" {
		t.Fatalf("Set failed: %v, %s", err, string(res))
	}

	// Get the value
	txGet := KVTx{Type: TxTypeGet, Key: "foo"}
	b, _ = json.Marshal(txGet)
	res, err = vm.ExecuteTx(ctx, b)
	if err != nil || string(res) != "bar" {
		t.Fatalf("Get failed: %v, %s", err, string(res))
	}
}

func TestKVStoreGetMissingKey(t *testing.T) {
	ctx := context.Background()
	vm := &EVM{}
	txGet := KVTx{Type: TxTypeGet, Key: "missing"}
	b, _ := json.Marshal(txGet)
	_, err := vm.ExecuteTx(ctx, b)
	if err == nil {
		t.Fatal("Expected error for missing key, got nil")
	}
}

func TestKVStoreInvalidTx(t *testing.T) {
	ctx := context.Background()
	vm := &EVM{}
	_, err := vm.ExecuteTx(ctx, []byte("not-json"))
	if err == nil {
		t.Fatal("Expected error for invalid tx, got nil")
	}
} 