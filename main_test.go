package main

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestIngestAndReplay(t *testing.T) {
	root, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}

	bundle := filepath.Join(t.TempDir(), "incident.bundle.json")
	ingestCmd := exec.Command("go", "run", ".", "ingest", filepath.Join(root, "examples", "incident.json"), "--output", bundle)
	ingestCmd.Dir = root
	ingestOutput, err := ingestCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("ingest failed: %v\n%s", err, string(ingestOutput))
	}

	replayCmd := exec.Command("go", "run", ".", "replay", bundle, "--candidate", filepath.Join(root, "examples", "fixed_incident.json"))
	replayCmd.Dir = root
	replayOutput, err := replayCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("replay failed: %v\n%s", err, string(replayOutput))
	}

	output := string(replayOutput)
	if !strings.Contains(output, "diff:") {
		t.Fatalf("expected diff output, got:\n%s", output)
	}
	if !strings.Contains(output, "$.actual_response.status") {
		t.Fatalf("expected status diff, got:\n%s", output)
	}
}
