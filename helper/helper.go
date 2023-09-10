package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, tickets uint, remainingTickets uint) (bool, bool, bool) {
	validUserName := len(firstName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@")
	validNumberTickets := remainingTickets >= tickets && tickets > 0
	return validUserName, validEmail, validNumberTickets
}
