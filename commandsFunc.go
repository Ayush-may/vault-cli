package main

func commandsFunc(commands []string) {
	if len(commands) < 2 {
		panic("Command length is lower!")
	}

	valtCommand := commands[1]
	orderCommand := commands[2]
	if valtCommand != "vault" {
		panic("Do you mean vault ?")
	}

	switch orderCommand {
	case "help":
		handleHelpCommand()
	case "new":
		handleNewCommand()
	case "list":
		handleListCommand()
	case "change-username":
		if len(commands) < 5 {
			panic("Not enough arguments for change-username!")
		}

		preUsername := commands[3]
		newUsername := commands[4]

		handleChangeUsernameCommand(preUsername, newUsername)
	case "change-password":
		if len(commands) < 5 {
			panic("Not enough arguments for change-username!")
		}

		username := commands[3]
		newPassword := commands[4]

		handleChangePasswordCommand(username, newPassword)
	case "delete":
		if len(commands) < 4 {
			panic("Not enough arguments for change-username!")
		}

		username := commands[3]
		handleUserDeleteCommand(username)
	default:
		panic("No command is matches!")

	}
}
