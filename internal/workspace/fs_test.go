package workspace

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "fs_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	testFile := filepath.Join(tmpdir, "test.txt")
	
	if FileExists(testFile) {
		t.Errorf("FileExists(%q) = true, want false", testFile)
	}

	if err := os.WriteFile(testFile, []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}

	if !FileExists(testFile) {
		t.Errorf("FileExists(%q) = false, want true", testFile)
	}
}

func TestEnsureDirectoryExists(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "fs_test_dir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	testDir := filepath.Join(tmpdir, "subdir", "deep")
	
	if _, err := os.Stat(testDir); !os.IsNotExist(err) {
		t.Fatalf("test directory already exists: %s", testDir)
	}

	if err := EnsureDirectoryExists(testDir); err != nil {
		t.Fatalf("EnsureDirectoryExists(%q) returned error: %v", testDir, err)
	}

	info, err := os.Stat(testDir)
	if err != nil {
		t.Fatalf("os.Stat(%q) returned error: %v", testDir, err)
	}

	if !info.IsDir() {
		t.Errorf("%q is not a directory", testDir)
	}
}
