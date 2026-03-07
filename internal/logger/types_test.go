package logger

import (
	"bytes"
	"testing"
)

func TestLogLevelComparison(t *testing.T) {
	if DEBUG >= INFO {
		t.Error("DEBUG should be less than INFO")
	}
	if INFO >= WARN {
		t.Error("INFO should be less than WARN")
	}
	if WARN >= ERROR {
		t.Error("WARN should be less than ERROR")
	}
}

func TestStructuredLoggerCreation(t *testing.T) {
	var buf bytes.Buffer
	l := &StructuredLogger{
		Writer:    &buf,
		Component: "test",
		PID:       1234,
	}

	if l.Component != "test" {
		t.Errorf("Expected component 'test', got %s", l.Component)
	}
	if l.PID != 1234 {
		t.Errorf("Expected PID 1234, got %d", l.PID)
	}
}
