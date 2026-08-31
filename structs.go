package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var fileName string = "vault.json"

type Profile struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// returns all profiles from vault.json, or creates the file if it doesn't exist
func returnAllProfilesFromJson() []Profile {
	var profiles []Profile

	data, err := os.ReadFile(fileName)
	if err != nil {
		// If file doesn't exist, create it with an empty json array []
		if os.IsNotExist(err) {
			_ = os.WriteFile(fileName, []byte("[]"), 0644)
			return profiles
		}
		fmt.Println("Error reading file:", err)
		return profiles
	}

	if len(data) == 0 {
		return profiles
	}

	err = json.Unmarshal(data, &profiles)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return profiles
	}

	return profiles
}

// saves full profiles slice into vault.json
func saveFullProfileIntoJson(profiles []Profile) {
	data, err := json.MarshalIndent(profiles, "", "    ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

// saves a new profile to vault.json
func (newProfile Profile) saveProfile() {
	allProfile := returnAllProfilesFromJson()

	// validating if same username already exists
	for _, profile := range allProfile {
		if profile.Username == newProfile.Username {
			fmt.Println("Username already exists!")
			return
		}
	}

	allProfile = append(allProfile, newProfile)
	saveFullProfileIntoJson(allProfile)
	fmt.Println("Successfully added!")
}
