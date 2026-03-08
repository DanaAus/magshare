//go:build !windows

package workspace

import "errors"

// CreateDesktopShortcut is a stub for non-Windows platforms.
func CreateDesktopShortcut() error {
	return errors.New("desktop shortcut creation is only supported on Windows")
}
