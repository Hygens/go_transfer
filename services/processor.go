package services

import (
	. "./go_transfer/models"
)

func ProcessTransfer(users *Users, senderName string,
	avail float64, targetName string, amount float64) (bool, User) {
	if (avail == float64(0) || amount == float64(0) || amount > avail) ||
		(senderName == "default" || targetName == "default") {
		return false, User{}
	}
	balance := avail - amount
	var ret User
	i := 0
	for k, user := range users.Users {
		if user.Name == senderName {
			users.Users[k].Balance = balance
			i++
		}
		if user.Name == targetName {
			users.Users[k].Balance += amount
			ret = users.Users[k]
			i++
		}
		if i == 2 {
			break
		}
	}
	SaveUsers(users)
	return true, ret
}
