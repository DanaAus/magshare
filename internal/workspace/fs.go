package workspace

import (
	"os"
)

// FileExists returns true if the specified file or directory exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// EnsureDirectoryExists creates the specified directory and any missing parent directories.
func EnsureDirectoryExists(path string) error {
	if FileExists(path) {
		return nil
	}
	return os.MkdirAll(path, 0755)
}
