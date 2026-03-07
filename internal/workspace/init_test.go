package workspace

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitializeWorkspaceAtPath(t *testing.T) {
	tmpBase, err := os.MkdirTemp("", "magshare_init_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpBase)

	testRoot := filepath.Join(tmpBase, "magshare_test")
	
	// 1. Initial run: should create folders and return isFirstRun = true
	isFirstRun, err := InitializeWorkspaceAtPath(testRoot)
	if err != nil {
		t.Fatalf("InitializeWorkspaceAtPath returned error: %v", err)
	}
	if !isFirstRun {
		t.Error("InitializeWorkspaceAtPath returned isFirstRun = false, want true")
	}

	// Verify structure
	if !FileExists(testRoot) {
		t.Error("Workspace root not created")
	}
	logsDir := filepath.Join(testRoot, "logs")
	if !FileExists(logsDir) {
		t.Error("Logs directory not created")
	}
	configPath := filepath.Join(testRoot, "config.json")
	if !FileExists(configPath) {
		t.Error("Config file not created")
	}

	// 2. Second run: should return isFirstRun = false
	isFirstRun, err = InitializeWorkspaceAtPath(testRoot)
	if err != nil {
		t.Fatalf("InitializeWorkspaceAtPath (second run) returned error: %v", err)
	}
	if isFirstRun {
		t.Error("InitializeWorkspaceAtPath (second run) returned isFirstRun = true, want false")
	}
}

func TestInitializeWorkspace(t *testing.T) {
	// This will use the actual system UserConfigDir/LOCALAPPDATA
	// We should be careful not to mess up actual config.
	// But InitializeWorkspace just checks if it exists.
	
	// We can at least call it and check it doesn't error.
	_, err := InitializeWorkspace()
	if err != nil {
		t.Errorf("InitializeWorkspace() returned error: %v", err)
	}
}
