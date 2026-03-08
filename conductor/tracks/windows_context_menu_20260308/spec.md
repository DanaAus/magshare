# Specification: Windows Context Menu Integration (Magshare)

## Overview
This track implements a seamless integration between Magshare and the Windows shell by adding a "Share via Magshare" option to the right-click context menu for all files and directories.

## Functional Requirements
1.  **Registry Registration (`AddToContextMenu`):**
    - Use the `golang.org/x/sys/windows/registry` package for all registry operations.
    - Register the context menu for the **Current User (HKCU)** to avoid requiring Administrator privileges.
    - Target **All Files**: Create keys under `HKEY_CURRENT_USER\Software\Classes\*\shell\Magshare`.
    - Target **Directories**: Create keys under `HKEY_CURRENT_USER\Software\Classes\Directory\shell\Magshare`.
    - **Menu Label:** Set the default value to "Share via Magshare".
    - **Execution Command:** Set the `command` subkey default value to `"{absolute_path_to_exe}" send "%1"`.
2.  **Registry Removal (`RemoveFromContextMenu`):**
    - Implement a clean removal function that deletes the created registry keys (including the `command` subkey).
    - Ensure it handles cases where the keys might already be missing.
3.  **Path Resolution:**
    - Automatically determine the absolute path of the currently running Magshare executable using `os.Executable()`.
4.  **Error Handling:**
    - Log any registry access errors to the diagnostic logs (INFO/ERROR level).
    - Fail silently for the end-user if the operation fails (e.g., during onboarding) to maintain a frictionless experience.

## Non-Functional Requirements
- **Standard Go Library:** Avoid external CLI wrapper dependencies (e.g., calling `reg.exe` via `os/exec`).
- **Build Constraints:** Use `//go:build windows` to ensure the code only compiles on Windows.
- **Stub Implementation:** Provide empty/no-op stubs for non-Windows platforms in `internal/workspace`.

## Acceptance Criteria
- [ ] Right-clicking any file shows "Share via Magshare".
- [ ] Right-clicking any directory shows "Share via Magshare".
- [ ] Clicking the option launches Magshare with the correct arguments (e.g., `magshare.exe send "C:\path\to\target"`).
- [ ] Running the removal function completely cleans up the registry entries.
- [ ] The feature does not require elevation to Administrator.

## Out of Scope
- Integration with macOS (Finder) or Linux (Nautilus/Dolphin) context menus in this track.
- Background context menu (right-clicking inside an empty folder).
- Custom menu icons.
