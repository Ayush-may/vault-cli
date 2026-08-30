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
	case "new":
		handleNewCommand()
	case "list":
		handleListCommand()
	case "change-username":
		preUsername := commands[3]
		newUsername := commands[4]
		handleChangeUsernameCommand(preUsername, newUsername)
	default:
		panic("No command is matches!")

	}
}
