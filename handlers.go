package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// ####################################
// COMMANDS FUNCTIONS START
// ####################################

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

	for idx, profile := range Profiles {

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

// ####################################
// HELPER FUNCTIONS END
// ####################################
