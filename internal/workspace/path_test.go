package workspace

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetWorkspaceRoot(t *testing.T) {
	root, err := GetWorkspaceRoot()
	if err != nil {
		t.Fatalf("GetWorkspaceRoot() returned error: %v", err)
	}

	if root == "" {
		t.Error("GetWorkspaceRoot() returned empty string")
	}

	// Verify it contains "magshare"
	if !strings.Contains(strings.ToLower(root), "magshare") {
		t.Errorf("GetWorkspaceRoot() path %q does not contain 'magshare'", root)
	}

	// Check if it's an absolute path
	if !filepath.IsAbs(root) {
		t.Errorf("GetWorkspaceRoot() path %q is not absolute", root)
	}

	// On Windows, it should be in LocalAppData
	if os.Getenv("LOCALAPPDATA") != "" {
		localAppData := os.Getenv("LOCALAPPDATA")
		if !strings.HasPrefix(strings.ToLower(root), strings.ToLower(localAppData)) {
			t.Errorf("GetWorkspaceRoot() path %q should be inside LocalAppData %q", root, localAppData)
		}
	}
}

func TestGetLogsDir(t *testing.T) {
	logsDir, err := GetLogsDir()
	if err != nil {
		t.Fatalf("GetLogsDir() returned error: %v", err)
	}

	root, _ := GetWorkspaceRoot()
	expected := filepath.Join(root, "logs")

	if logsDir != expected {
		t.Errorf("GetLogsDir() = %q, want %q", logsDir, expected)
	}
}
