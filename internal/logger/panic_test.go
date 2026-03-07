package logger

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHandlePanic(t *testing.T) {
	tmpdir, err := os.MkdirTemp("", "panic_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	logPath := filepath.Join(tmpdir, "crash.log")
	
	// Create an empty log file first
	if err := os.WriteFile(logPath, []byte(""), 0644); err != nil {
		t.Fatal(err)
	}

	// We'll simulate a panic by calling HandlePanic manually in a way that it sees a recovery
	// This is hard because recover() only works inside a defer.
	// So we'll wrap it in a function.
	
	finished := make(chan bool)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// This is what HandlePanic should do
				HandlePanic(logPath, r)
				finished <- true
			}
		}()
		panic("test panic")
	}()

	<-finished

	// Verify log file contains the panic message
	content, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(content), "test panic") {
		t.Errorf("Log file does not contain panic message. Content: %s", string(content))
	}
	
	if !strings.Contains(string(content), "STACK TRACE") {
		t.Error("Log file does not contain stack trace")
	}
}
