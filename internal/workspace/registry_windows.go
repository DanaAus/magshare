//go:build windows

package workspace

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

// RegisterContextMenu adds 'Share via Magshare' to the Windows right-click menu for files and directories.
func RegisterContextMenu() error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}
	exePath, err = filepath.Abs(exePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Command: "C:\path\to\magshare.exe" send "%1"
	command := fmt.Sprintf("\"%s\" send \"%%1\"", exePath)
	const menuTitle = "Share via Magshare"

	targets := []string{
		`Software\Classes\*\shell\Magshare`,
		`Software\Classes\Directory\shell\Magshare`,
	}

	for _, target := range targets {
		// Create/Open the shell key
		key, _, err := registry.CreateKey(registry.CURRENT_USER, target, registry.ALL_ACCESS)
		if err != nil {
			return fmt.Errorf("failed to create registry key %s: %w", target, err)
		}

		if err := key.SetStringValue("", menuTitle); err != nil {
			key.Close()
			return fmt.Errorf("failed to set menu title for %s: %w", target, err)
		}

		// Create/Open the command subkey
		cmdKey, _, err := registry.CreateKey(key, "command", registry.ALL_ACCESS)
		key.Close() // Close parent
		if err != nil {
			return fmt.Errorf("failed to create command key for %s: %w", target, err)
		}

		if err := cmdKey.SetStringValue("", command); err != nil {
			cmdKey.Close()
			return fmt.Errorf("failed to set command for %s: %w", target, err)
		}
		cmdKey.Close()
	}

	return nil
}

// UnregisterContextMenu removes Magshare from the Windows right-click menu.
func UnregisterContextMenu() error {
	targets := []string{
		`Software\Classes\*\shell\Magshare`,
		`Software\Classes\Directory\shell\Magshare`,
	}

	for _, target := range targets {
		// First, delete the "command" subkey
		_ = registry.DeleteKey(registry.CURRENT_USER, target+`\command`)

		// Now delete the Magshare key itself
		lastSlash := strings.LastIndex(target, `\`)
		if lastSlash == -1 {
			continue
		}
		parentPath := target[:lastSlash]
		childName := target[lastSlash+1:]

		parentKey, err := registry.OpenKey(registry.CURRENT_USER, parentPath, registry.ALL_ACCESS)
		if err != nil {
			continue
		}
		_ = registry.DeleteKey(parentKey, childName)
		parentKey.Close()
	}

	return nil
}
