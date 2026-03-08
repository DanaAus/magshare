# Specification: Fix Context Menu Launch Error

## Overview
This track fixes a bug where Magshare fails to launch correctly when triggered via the Windows right-click context menu. Instead of performing the "send" action or showing the menu, the application displays a guard message: "This is a command line tool..." and exits.

## Functional Requirements
1.  **Remove or Refine Execution Guard:**
    - Identify and refactor the code in `main.go` or `cmd/` that prevents the application from running when launched via the shell context menu.
    - Ensure that valid command-line arguments (like `send "C:\path\to\file"`) are honored even if the program detects it's not being run from an existing interactive terminal session.
2.  **Context Menu Behavior:**
    - When launched via "Share via Magshare", the application should NOT just show the guard message.
    - It should ideally enter the **Prefilled Menu** state: launch the interactive mode with the target path already set, or proceed directly to the "Send" server setup while keeping the terminal window open.
3.  **Terminal Window Stability:**
    - Ensure the terminal window spawned by Windows Explorer remains open during the sharing session so the user can see the QR code.

## Non-Functional Requirements
- **Robustness:** If the arguments are malformed, provide a helpful error message instead of a generic guard.
- **Platform Specificity:** The fix should primarily target Windows behavior while remaining cross-platform compatible.

## Acceptance Criteria
- [ ] Clicking "Share via Magshare" on a file/folder no longer displays the "This is a command line tool..." error.
- [ ] The application opens a terminal window and correctly processes the `send` command for the target file.
- [ ] The interactive loop (if triggered) allows for further actions or a graceful exit.

## Out of Scope
- Redesigning the entire CLI argument parser.
- Adding macOS/Linux context menu support in this fix.
