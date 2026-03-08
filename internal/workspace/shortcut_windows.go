//go:build windows

package workspace

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// CreateDesktopShortcut creates a shortcut to Magshare on the user's desktop.
func CreateDesktopShortcut() error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}
	exePath, err = filepath.Abs(exePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	desktopPath, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	desktopPath = filepath.Join(desktopPath, "Desktop")

	shortcutPath := filepath.Join(desktopPath, "Magshare.lnk")

	// Use PowerShell to create the shortcut via COM (WScript.Shell)
	psCommand := fmt.Sprintf(
		"$s=(New-Object -ComObject WScript.Shell).CreateShortcut('%s');$s.TargetPath='%s';$s.WorkingDirectory='%s';$s.Description='Magshare File Sharing';$s.Save()",
		shortcutPath, exePath, filepath.Dir(exePath),
	)

	cmd := exec.Command("powershell", "-NoProfile", "-Command", psCommand)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("powershell failed: %w (output: %s)", err, string(output))
	}

	return nil
}
