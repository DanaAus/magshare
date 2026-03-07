package logger

import (
	"io"
)

// LogLevel defines the severity of a log message.
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// String returns the string representation of the LogLevel.
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Logger is the interface for logging messages.
type Logger interface {
	Log(level LogLevel, msg string)
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

// StructuredLogger implements the Logger interface with formatted output.
type StructuredLogger struct {
	Writer    io.Writer
	Component string
	PID       int
}

// Log writes a message at the specified level.
func (l *StructuredLogger) Log(level LogLevel, msg string) {
	// Implementation will follow in formatting task
}

func (l *StructuredLogger) Debug(msg string) { l.Log(DEBUG, msg) }
func (l *StructuredLogger) Info(msg string)  { l.Log(INFO, msg) }
func (l *StructuredLogger) Warn(msg string)  { l.Log(WARN, msg) }
func (l *StructuredLogger) Error(msg string) { l.Log(ERROR, msg) }
