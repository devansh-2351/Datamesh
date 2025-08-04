package consensus

import (
	"context"
	"strings"
	"testing"
)

func TestConsensusEngineLifecycle(t *testing.T) {
	ctx := context.Background()
	engine := &TendermintEngine{}

	if err := engine.Start(ctx); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	status := engine.Status()
	if !strings.Contains(status, "running: true") {
		t.Fatalf("Expected running: true, got: %s", status)
	}

	engine.IncrementBlock()
	engine.IncrementBlock()
	status = engine.Status()
	if !strings.Contains(status, "block height: 2") {
		t.Fatalf("Expected block height: 2, got: %s", status)
	}

	if err := engine.Stop(ctx); err != nil {
		t.Fatalf("Stop failed: %v", err)
	}
	status = engine.Status()
	if !strings.Contains(status, "running: false") {
		t.Fatalf("Expected running: false, got: %s", status)
	}
} 