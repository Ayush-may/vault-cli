package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// ####################################
// COMMANDS FUNCTIONS START
// ####################################

func handleHelpCommand() {
	fmt.Println(`
Valt - Simple Password Vault

Usage:
   vault new
   vault list
   vault get <username>
   vault update <username>
   vault delete <username>
   vault change-username <old> <new>
   vault change-password <username> <new-password>
     `)
}

// vault new
func handleNewCommand() {
	var user Profile

	fmt.Print("Enter Name : ")
	fmt.Scanln(&user.Name)

	fmt.Print("Enter Username : ")
	fmt.Scanln(&user.Username)

	fmt.Print("Enter password : ")
	fmt.Scanln(&user.Password)

	user.saveProfile()
	fmt.Println("Successfully added!")
}

// vault list
func handleListCommand() {
	profiles := returnAllProfilesFromJson()

	for idx, profile := range profiles {
		fmt.Printf("index : %d \nName : %s \nUsername : %s \npassword : %s \n\n", idx, profile.Name, profile.Username, profile.Password)
	}

}

// vault change-username existing_username new_username
func handleChangeUsernameCommand(preUsername, newUsername string) {
	Profiles := returnAllProfilesFromJson()
	isExist := false

	for idx, profile := range Profiles {
		if profile.Username == preUsername {
			Profiles[idx].Username = newUsername
			saveFullProfileIntoJson(Profiles)

			fmt.Println("Username is successfully updated now!")
			isExist = true
			break
		}
	}

	if !isExist {
		fmt.Println("This username isnt exist!")
	}
}

// vault change-username existing_username new_username
func handleChangePasswordCommand(username, newPassword string) {
	Profiles := returnAllProfilesFromJson()
	isExist := false

	for idx, profile := range Profiles {
		if profile.Username == username {
			Profiles[idx].Password = newPassword
			saveFullProfileIntoJson(Profiles)

			fmt.Println("Password is successfully updated now!")
			isExist = true
			break
		}
	}

	if !isExist {
		fmt.Println("This username isnt exist!")
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

// ####################################
// HELPER FUNCTIONS START
// ####################################

// this returns all the profiles from the json file
// file name : vault.json
func returnAllProfilesFromJson() []Profile {
	var profile []Profile

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic("Unable to read file in method returnAllProfilesFromJson()")
	}

	err = json.Unmarshal(data, &profile)
	if err != nil {
		panic("Unable to unmarsha in returnAllProfilesFromJson()")
	}

	return profile
}

// this saves profiles into the json
func saveFullProfileIntoJson(profiles []Profile) {
	data, err := json.MarshalIndent(profiles, "", " ")
	if err != nil {
		panic("Error while converting profiles into json!")
	}

	err = os.WriteFile(fileName, data, 0600)
	if err != nil {
		panic("Error while writing on file!")
	}
}

// ####################################
// HELPER FUNCTIONS END
// ####################################
