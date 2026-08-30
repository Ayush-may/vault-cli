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

// saves data to the vault.json
func (newProfile Profile) saveProfile() {
	var allProfile []Profile

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		panic("Unable to read the file")
	}

	err = json.Unmarshal(data, &allProfile)
	if err != nil {
		fmt.Println(err)
		panic("Unable to perform json unmarshal!")
	}

	// validating if same username is already exist or not
	for _, profile := range allProfile {
		if profile.Username == newProfile.Username {
			panic("Username is already exist!")
		}
	}

	allProfile = append(allProfile, newProfile)

	data, err = json.Marshal(allProfile)
	if err != nil {
		fmt.Println(err)
		panic("Unable to perform json marshal!")
	}

	// writing the file with new data here
	err = os.WriteFile(fileName, data, 0600)
}
