package services

import (
	. "github.com/Hygens/go_transfer/models"
)

func ProcessTransfer(users *Users, senderName string,
	avail float64, targetName string, amount float64) (bool, *User) {
	if (avail == float64(0) || amount == float64(0) || amount > avail) ||
		(senderName == "default" || targetName == "default") ||
		(senderName == targetName) {
		return false, &User{}
	}
	balance := avail - amount
	var receiver *User
	var k1, k2 int
	k1, _ = GetUser(senderName, users)
	users.Users[k1].Balance = balance
	k2, receiver = GetUser(targetName, users)
	receiver.Balance += amount
	users.Users[k2].Balance = receiver.Balance
	SaveUsers(users)
	return true, receiver
}
