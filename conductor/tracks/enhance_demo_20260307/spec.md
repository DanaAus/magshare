# Track Specification: Enhance Demo Mode Customization

## Overview
This track aims to enhance the existing "Demo Mode" in magshare to allow users to customize key parameters like the network port and the security PIN, bringing it closer to the functionality of the standard "send" and "receive" modes.

## Functional Requirements
- **Port Customization:** Allow users to specify a custom port for the demo server using a flag (e.g., `--port`).
- **PIN Customization:** Allow users to specify a custom 4-digit PIN for the demo session using a flag (e.g., `--pin`).
- **Validation:** Ensure that the provided port is a valid, available network port and the PIN is exactly 4 digits.

## Technical Details
- Update the demo command definition in `cmd/demo.go` (or wherever the demo mode is defined) to include the new flags.
- Pass the custom port and PIN values to the internal server initialization logic.
- Ensure the TUI (if applicable) also reflects these custom settings.

## Acceptance Criteria
- User can run `magshare demo --port 8080` and the server starts on port 8080.
- User can run `magshare demo --pin 1234` and the security PIN is set to 1234.
- User can combine both flags.
- Default behavior remains unchanged if flags are not provided.
