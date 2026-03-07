# Product Guidelines: magshare

## Tone & Style
The tone of magshare should be **Friendly & Accessible**. Documentation, CLI messages, and UI text should be helpful, welcoming, and easy to understand for both technical and non-technical users. Avoid overly dense technical jargon when simpler explanations suffice.

## User Experience (UX) Principles
*   **Minimal Friction:** Every action should require the fewest possible steps. Default to the most common use case (e.g., auto-detecting IP, choosing a high-range port).
*   **Visual Clarity & Feedback:** Use clear, consistent formatting in the terminal. Provide immediate and informative feedback for every user action (e.g., "Server started at...", "File received: ...").
*   **Convention over Configuration:** Provide sensible defaults that work for 90% of users, while allowing advanced users to override settings via flags or the TUI.

## Accessibility (A11y)
*   **High Contrast Terminal UI:** Ensure that TUI elements (using Lipgloss/Bubbletea) use colors that are legible across various terminal themes (light and dark).
*   **Web Accessibility (A11y):** The "Receive Dropzone" web UI must follow basic accessibility guidelines, including semantic HTML and proper ARIA labels.
*   **Keyboard Accessible UI:** Both the TUI and the web-based Dropzone must be fully navigable using only a keyboard.

## Design Guidelines
*   **Terminal UI (TUI):** Utilize the Charmbracelet stack (Bubbletea, Huh, Lipgloss) for a modern, responsive, and visually appealing terminal experience.
*   **Web UI:** The Dropzone should be a single-page, responsive HTML5 interface that is lightweight and fast-loading.
