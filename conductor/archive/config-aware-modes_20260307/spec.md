# Specification: Config-Aware Application Modes

## Overview
Ensure all application modes (send, receive, and interactive) correctly utilize the settings defined during the first-time setup process. This involves loading the `config.json` from the Magshare workspace and using its values as defaults for all operations, while still allowing manual overrides via CLI flags.

## Functional Requirements
1.  **Configuration Loading:**
    -   Implement a robust `LoadConfig()` function to read and parse `config.json`.
    -   Load the configuration once during the root command's initialization.
2.  **Global Configuration Context:**
    -   Store the loaded configuration in a globally accessible location within the `cmd` package.
3.  **Command Integration:**
    -   **`receive` mode:**
        -   Use `DownloadDir` from the config as the default destination for received files.
        -   Use `SecureMode` from the config as the default security setting.
    -   **`send` mode:**
        -   Use `SecureMode` from the config as the default security setting.
    -   **`interactive` mode:**
        -   Pre-populate interactive prompts with values from the config.
4.  **Precedence Logic:**
    -   The application must strictly follow this order of priority:
        1.  Explicit CLI flags (e.g., `--port`, `--secure`).
        2.  Values from the `config.json` file.
        3.  Hardcoded application defaults.
5.  **Robust Validation:**
    -   Verify the existence of the `DownloadDir` when loading the config.
    -   Ensure the `Port` is within the valid range (1-65535).

## Non-Functional Requirements
-   **Consistency:** All commands should behave predictably regardless of whether they are run directly or via the interactive menu.
-   **Error Handling:** If the config file is corrupted or unreadable, the application should log a warning and fall back to hardcoded defaults rather than crashing.

## Acceptance Criteria
- [ ] `magshare receive` automatically saves files to the directory specified during setup.
- [ ] `magshare send` and `magshare receive` default to the security mode (PIN) chosen during setup if no flag is provided.
- [ ] Explicit CLI flags (like `--secure=false`) correctly override the config file's "always secure" setting.
- [ ] Corrupted config files do not prevent the app from starting (graceful fallback).

## Out of Scope
-   **Log Rotation/Pruning:** Managing the lifecycle of multiple historical logs.
-   **Hot-Reloading:** Re-loading the config file if it changes while the server is running.
