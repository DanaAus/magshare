package logger

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

// HandlePanic processes a recovered panic by logging it and waiting before exit.
// r is the value returned by recover().
func HandlePanic(logPath string, r interface{}) {
	if r == nil {
		return
	}

	// Capture stack trace
	stack := debug.Stack()

	// Prepare crash report
	report := fmt.Sprintf("\n========== CRASH REPORT ==========\n")
	report += fmt.Sprintf("Time: %s\n", time.Now().Format(time.RFC3339))
	report += fmt.Sprintf("Panic: %v\n\n", r)
	report += "STACK TRACE:\n"
	report += string(stack)
	report += "==================================\n"

	// Write to terminal
	fmt.Fprintf(os.Stderr, "\n[CRITICAL ERROR] magshare has crashed!\n")
	fmt.Fprintf(os.Stderr, "Panic: %v\n", r)
	fmt.Fprintf(os.Stderr, "Crash log saved to: %s\n", logPath)

	// Write to log file
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		f.WriteString(report)
		f.Close()
	}

	// Countdown timer for 5 seconds
	fmt.Fprintf(os.Stderr, "\nClosing in ")
	for i := 5; i > 0; i-- {
		fmt.Fprintf(os.Stderr, "%d... ", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintln(os.Stderr)
}
