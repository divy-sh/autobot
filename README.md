# Autobot 🚀

**Autobot** is a lightweight CLI tool built with **Golang** that keeps your device active and maintains your team’s availability status automatically.

---

## Features
- Prevents your system from going idle.
- Keeps your team status active.
- Simple command-line interface.

---

## Installation
Make sure you have [Go](https://golang.org/dl/) installed, then install **Autobot**:

```bash
go install github.com/divy-sh/autobot@latest
```

---

## Usage
Run the program in your terminal:

```bash
autobot -seconds=<interval>
```

- `<interval>` (optional) — Time interval in seconds to reset the sleep timer. Default behavior is automatic if not specified.

**Example:**

```bash
autobot -seconds=60
```

This will reset the sleep timer every 60 seconds.

---

## License
MIT License
