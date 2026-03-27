<p align="center">
  <img src="assets/images/hero.png" alt="freaky_md WhatsApp Bot Hero Image" width="800">
</p>

# freaky-md

A lightweight, modular WhatsApp MD Bot written in Go. No bloat, no unnecessary abstractions—just a clean command handler backed by SQLite.

## Features
- **Fast:** Built on `whatsmeow` for high-performance WhatsApp Web protocol interaction.
- **Persistent:** Uses SQLite for session management and data storage.
- **Modular:** Simple command registry; adding new functionality takes seconds.
- **Simple Login:** Generates a QR code directly in your terminal.

## Tech Stack
- **Language:** Go 1.26.1
- **Protocol:** [whatsmeow](https://github.com/tulir/whatsmeow)
- **Database:** SQLite3

## Getting Started

### Prerequisites
- Go 1.26 or higher.
- GCC (for `go-sqlite3` compilation).

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/zaid-nizam/freaky-md.git
   cd freaky-md
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the bot:
   ```bash
   go run cmd/bot/main.go
   ```

### Authentication
On the first run, the bot will display a QR code in the terminal. Scan it with your WhatsApp mobile app (**Linked Devices > Link a Device**) to authenticate. Your session will be saved in `freaky.db`.

## Commands
The default prefix is `!`.
- `!ping` - Check if the bot is alive.

## Development
To add a new command, create a file in `internal/commands/` and register it in the `Registry`.
