package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name    string  `json:"name"`
	Account string  `json:"account"`
	Balance float64 `json:"balance"`
}

func GetUsers() Users {
	byteValue, err := ioutil.ReadFile("models/data.json")

	if err != nil {
		fmt.Println("File reading error", err)
	}
	fmt.Println("Successfully Opened data.json")

	var users Users
	json.Unmarshal(byteValue, &users)

	return users
}

func SaveUsers(users *Users) {
	byteValue, _ := json.MarshalIndent(users, "", "")
	_ = ioutil.WriteFile("models/data.json", byteValue, 0644)
}
