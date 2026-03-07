<div align="center">

# 🚀 Magshare

[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/DanaAus/magshare)](https://goreportcard.com/report/github.com/DanaAus/magshare)

**Instant, frictionless file sharing and receiving across your local network via CLI.**

<br>

<img src="data:image/gif;base64,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y4sOHDiBMrXsy4sePHkCNLnky5suXLmDNr3sy5s+fPoEOLHk26tOnTqFOrXs26tevXsGPLnk27tu3buB/SzF10N2+jAQEAIfkEBQIA+QAsAAAAAAEAAQAACAQA8wUEACH5BAUDAPkALJsBdAIBAAEAAAYDwFMQACH5BAkCAPkALEIBqwHhABEBAAj/APNFyUewoMGDCBMqXMiwocOHECNKnEixokWINC5q3Mixo8ePIEMODEmypMmTKE2WSMmypcuXMEfCnEmzps2bOHPq3Mmzp8+fQIMKHUq0qNGjSJMqXcq0qdOnUKNKnUq1qtWrWLNq3cq1q9evYMOKHUu2rNmzaNOqXcu2rdu3cOPKnUu3rt27ePPq3cu3r9+/gAMLHky4sOHDiBMrXsy4sePHkCNLnky5suXLmDNr3sy5s+fPoEOLHk26tOnTqFOrXs26tevXsGPLnk27tu3buHPr3s27t+/fwIMLH068uPHjyJMrX868ufPn0KNLn069uvXr2LNr3869u/fv4MOLpx9Pvrz58+jTq1/Pvr379/Djy59Pv779+/jz69/Pv7///wAGKOCABBZo4IEIJqjgggw26OCDEEYo4YQUVmjhhRhmqOGGHHbo4YcghijiiCSWaOKJKKao4oostujiizDGKOOMNNZo44045qjjjjz26OOPQAYp5JBEFmnkkUgmqeSSTDbp5JNQRinllFRWaeWVWGap5ZZcdunll2CGKeaYZJZp5pnrWREQADs=" alt="magshare demo" width="800">

</div>

---

**Magshare** is a blazing-fast, terminal-based utility designed to eliminate the friction of transferring files between devices on the same local network. By spinning up an ephemeral local web server and rendering a QR code directly in your terminal, it allows any mobile device or PC to securely download or upload files in seconds—no cables, no cloud drives, and no client-side app installations required.

## ✨ Core Features

*   📱 **Instant QR Generation:** Automatically detects your local IP and renders an access URL as a scannable QR code directly within your terminal window.
*   ⚡ **Interactive TUI Mode:** Don't want to remember command flags? Launch the guided, prompt-based UI to effortlessly configure your sharing session.
*   🪶 **Memory-Efficient Streaming:** Engineered to handle ultra-large files (10GB+) by streaming directly from disk, keeping RAM consumption safely under 20MB.
*   🛡️ **Secure Mode:** Protect sensitive network transfers by requiring a dynamically generated 4-digit PIN before a download or upload can begin.
*   🌐 **Web Dropzone:** Running in "receive" mode serves a lightweight, responsive HTML5 dropzone to the client device for seamless drag-and-drop uploads.

---

## 📦 Installation

Magshare is distributed as a standalone executable. Choose your preferred package manager below:

### Windows Native
**Using Scoop:**
```powershell
scoop bucket add magshare https://github.com/DanaAus/magshare
scoop install magshare
```

**Using WinGet:**
```powershell
winget install magshare
```

### Cross-Platform (Node / JavaScript Ecosystem)
If you already have a Node or Bun environment, you can install or run magshare globally:
```bash
bun x magshare
# OR
npm install -g magshare
```

### From Source (Go)
For developers who want to compile the latest version directly:
```bash
go install github.com/DanaAus/magshare@latest
```

---

## 🚀 Usage

magshare is designed to be completely intuitive. You can use the guided TUI, or pass commands directly.

### 1. Interactive TUI Mode
Simply run the command with no arguments to launch the interactive terminal interface:
```bash
magshare
```

### 2. Sending a File (Host ➔ Client)
Share a specific file. The client's browser will automatically prompt a direct download upon scanning the QR code.
```bash
magshare send ./my-file.txt
```
*Tip: Add the `--secure` flag to generate a one-time PIN for the transfer.*

### 3. Receiving Files (Client ➔ Host)
Spin up a temporary web server that allows anyone on the network to upload files directly to your current terminal directory.
```bash
magshare receive
```

---

## 🤝 Contributing
magshare is open-source software, and contributions are always welcome! If you'd like to improve the codebase, add a feature, or report a bug:
1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

## 📄 License
This project is distributed under the [Apache 2.0 License](LICENSE). Feel free to use, modify, and distribute it as you see fit.
