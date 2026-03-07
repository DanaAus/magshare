# Specification: Resumable & Secure Large File Streaming

## Overview
Optimize the file-serving capability of `magshare` to handle large files efficiently and reliably. This involves supporting byte-range requests (resume) for mobile compatibility, ensuring strictly bounded memory usage, and reinforcing path security.

## Functional Requirements
1.  **Resumable Downloads (Byte Range Support):**
    -   Integrate `http.ServeContent` to automatically handle `Range`, `If-Range`, and `Content-Range` headers.
    -   Ensure mobile browsers can pause and resume downloads without restarting from 0%.
2.  **Memory Optimization:**
    -   Guarantee RAM usage remains under 20MB regardless of file size by utilizing `io.ReadSeeker` and standard library streaming.
3.  **Enhanced Security:**
    -   **Path Sanitization:** Clean target paths to prevent any directory traversal attempts.
    -   **Boundary Check:** Verify that the served file is the intended target and not a system file accessed via illegal path manipulation.
4.  **Integrated Progress Feedback:**
    -   Implement a `ProgressReadSeeker` that wraps the file and updates the terminal progress bar.
    -   For range requests, the progress bar should represent the progress of the *requested chunk*.

## Non-Functional Requirements
-   **Compatibility:** Must work seamlessly with the existing ephemeral server and PIN/Hash security layer.
-   **Robustness:** Gracefully handle "Connection Reset" or "Broken Pipe" errors without crashing the main application.

## Acceptance Criteria
- [ ] Large files (e.g., 5GB+) can be downloaded with < 20MB RAM overhead.
- [ ] Resuming a download in Chrome/Safari/Mobile works (verified via `Range` header simulation).
- [ ] Path traversal attempts (e.g., `../../etc/passwd`) are blocked.
- [ ] Terminal progress bar correctly reflects the bytes sent during a specific request.

## Out of Scope
-   **Multi-threaded Downloads:** Optimizing for multiple concurrent range requests for the same file (standard single-stream behavior is prioritized).
-   **Disk Caching:** No intermediate server-side caching of file chunks.
