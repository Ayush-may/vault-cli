package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ####################################
// COMMANDS FUNCTIONS START
// ####################################

func handleHelpCommand() {
	fmt.Print(`
Valt - Simple Password Vault

Usage:
   vault help
   vault new
   vault list
   vault get <username>
   vault change-username <old-username> <new-username>
   vault change-password <username> <new-password>
   vault delete <username>
`)
}

// vault new
func handleNewCommand() {
	scanner := bufio.NewScanner(os.Stdin)
	var user Profile

	fmt.Print("Enter Name : ")
	if scanner.Scan() {
		user.Name = strings.TrimSpace(scanner.Text())
	}

	fmt.Print("Enter Username : ")
	if scanner.Scan() {
		user.Username = strings.TrimSpace(scanner.Text())
	}

	fmt.Print("Enter Password : ")
	if scanner.Scan() {
		user.Password = strings.TrimSpace(scanner.Text())
	}

	if user.Username == "" {
		fmt.Println("Username cannot be empty!")
		return
	}

	user.saveProfile()
}

// vault list
func handleListCommand() {
	profiles := returnAllProfilesFromJson()

	if len(profiles) == 0 {
		fmt.Println("No profiles found in vault.")
		return
	}

	for idx, profile := range profiles {
		fmt.Printf("index : %d \nName : %s \nUsername : %s \nPassword : %s \n\n", idx, profile.Name, profile.Username, profile.Password)
	}
}

// vault get <username>
func handleGetCommand(username string) {
	profiles := returnAllProfilesFromJson()

	for idx, profile := range profiles {
		if profile.Username == username {
			fmt.Printf("index : %d \nName : %s \nUsername : %s \nPassword : %s \n\n", idx, profile.Name, profile.Username, profile.Password)
			return
		}
	}

	fmt.Println("User not found!")
}

// vault change-username <old-username> <new-username>
func handleChangeUsernameCommand(preUsername, newUsername string) {
	profiles := returnAllProfilesFromJson()
	isExist := false

	// check if new username is already taken
	for _, profile := range profiles {
		if profile.Username == newUsername {
			fmt.Println("The new username is already taken!")
			return
		}
	}

	for idx, profile := range profiles {
		if profile.Username == preUsername {
			profiles[idx].Username = newUsername
			saveFullProfileIntoJson(profiles)
			fmt.Println("Username is successfully updated now!")
			isExist = true
			break
		}
	}

	if !isExist {
		fmt.Println("This username doesn't exist!")
	}
}

// vault change-password <username> <new-password>
func handleChangePasswordCommand(username, newPassword string) {
	profiles := returnAllProfilesFromJson()
	isExist := false

	for idx, profile := range profiles {
		if profile.Username == username {
			profiles[idx].Password = newPassword
			saveFullProfileIntoJson(profiles)
			fmt.Println("Password is successfully updated now!")
			isExist = true
			break
		}
	}

	if !isExist {
		fmt.Println("This username doesn't exist!")
	}
}

// vault delete <username>
func handleUserDeleteCommand(username string) {
	profiles := returnAllProfilesFromJson()
	isDeleted := false

	for idx, profile := range profiles {
		if profile.Username == username {
			profiles = append(profiles[:idx], profiles[idx+1:]...)
			saveFullProfileIntoJson(profiles)
			fmt.Println("User successfully deleted!")
			isDeleted = true
			break
		}
	}

	if !isDeleted {
		fmt.Println("Unable to find this user!")
	}
}

// ####################################
// COMMANDS FUNCTIONS END
// ####################################
