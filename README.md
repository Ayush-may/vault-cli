# Valt

A simple CLI password vault built with Go.

## How to Run

Run directly:
```bash
go run . <command>
```

Or build the binary:
```bash
go build -o vault.exe
./vault <command>
```

---

## Commands

| Command | Description |
| :--- | :--- |
| `vault help` | Show help and available commands |
| `vault new` | Add a new account (Name, Username, Password) |
| `vault list` | Show all saved accounts |
| `vault get <username>` | Get details of a specific account |
| `vault change-username <old-username> <new-username>` | Update an existing username |
| `vault change-password <username> <new-password>` | Update a password for a user |
| `vault delete <username>` | Delete an account from the vault |

---

## Data Storage

All data is stored locally in `vault.json`. The file is created automatically on first use.
