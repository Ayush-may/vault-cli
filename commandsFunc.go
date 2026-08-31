package main

import "fmt"

func commandsFunc(commands []string) {
	// If the user typed "vault" as the first argument (e.g. "go run . vault list"), skip it
	if len(commands) > 1 && commands[1] == "vault" {
		commands = append(commands[:1], commands[2:]...)
	}

	// If no subcommand is provided (e.g. just running "vault"), show help
	if len(commands) < 2 {
		handleHelpCommand()
		return
	}

	orderCommand := commands[1]

	switch orderCommand {
	case "help":
		handleHelpCommand()
	case "new":
		handleNewCommand()
	case "list":
		handleListCommand()
	case "get":
		if len(commands) < 3 {
			fmt.Println("Usage: vault get <username>")
			return
		}
		handleGetCommand(commands[2])
	case "change-username":
		if len(commands) < 4 {
			fmt.Println("Usage: vault change-username <old-username> <new-username>")
			return
		}
		preUsername := commands[2]
		newUsername := commands[3]
		handleChangeUsernameCommand(preUsername, newUsername)
	case "change-password":
		if len(commands) < 4 {
			fmt.Println("Usage: vault change-password <username> <new-password>")
			return
		}
		username := commands[2]
		newPassword := commands[3]
		handleChangePasswordCommand(username, newPassword)
	case "delete":
		if len(commands) < 3 {
			fmt.Println("Usage: vault delete <username>")
			return
		}
		username := commands[2]
		handleUserDeleteCommand(username)
	default:
		fmt.Printf("Unknown command: %s\nRun 'vault help' to see available commands.\n", orderCommand)
	}
}
