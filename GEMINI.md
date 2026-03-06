# QShare Project Overview

QShare is a terminal-based utility designed for instant, frictionless file sharing and receiving across a local network. It leverages ephemeral web servers and terminal-rendered QR codes to allow any device on the same Wi-Fi network to securely download or upload files without complex configuration.

## Core Technologies
- **Language:** Go (Golang) - chosen for its static compilation, network capabilities, and zero-dependency binaries.
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra) - used for command-line interface structure, flags, and help menus.
- **QR Rendering:** [qrterminal](https://github.com/mdp/qrterminal) - renders QR codes directly in the terminal using ANSI escape codes.
- **Networking:** Go's standard `net/http` and `net` packages for robust, secure server and interface discovery.
- **Embedded Assets:** Go's `embed` package to bundle the "Receive Dropzone" UI directly into the single executable.

## Project Structure
- `main.go`: Application entry point.
- `cmd/`: CLI command definitions (root, send, receive).
- `internal/`:
    - `handlers/`: High-level business logic for handling send and receive operations.
    - `server/`: Implementation of the `EphemeralServer`, which manages its own lifecycle and timeouts.
    - `network/`: Logic for network interface discovery and auto-selection.
- `ui/`: Contains the `dropzone.html` and the `embed.go` file to bundle frontend assets.
- `docs/`: Technical specifications and product documentation.

## Building and Running

### Prerequisites
- Go 1.25 or higher.

### Key Commands
- **Build:** `go build -o qshare.exe`
- **Run (Development):** `go run main.go [command]`
- **Test:** `go test ./...` (TODO: Implement automated tests)

### Example Usage
- **Send a file:** `qshare send presentation.pdf`
- **Send a directory (zipped on-the-fly):** `qshare send ./documents`
- **Receive files:** `qshare receive`
- **Secure Send (with PIN):** `qshare send file.txt --secure`

## Development Conventions

- **Architecture:** Follows the standard Go project layout with `cmd/` for CLI entry points and `internal/` for private application logic.
- **CLI Design:** All commands should be implemented using Cobra in the `cmd/` package. Use meaningful short and long descriptions.
- **Error Handling:** Errors should be handled gracefully. Commands typically use `log.Fatalf` for terminal errors, while internal logic should return `error` types for callers to handle.
- **Concurrency:** Uses Go routines and channels for managing the server lifecycle (e.g., `EphemeralServer.quitChan`).
- **Static Assets:** Any UI components must be placed in the `ui/` directory and registered in `ui/embed.go` to ensure they are compiled into the binary.
- **Security:** Prioritize randomized high-range ports and unpredictable URL paths for every sharing session.
