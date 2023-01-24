package helper

import (
	"strings"
)

// To export the fuction, the first char of the function must be capitalized.
func ValidateUserInput(
	firstName string,
	lastName string,
	email string,
	userTickets uint,
	remainingTickets uint,
) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	// isValidCity := city == "Singapore" || city == "London"

	// In Go, multiple values can be returned
	return isValidName, isValidEmail, isValidTicketNumber
}
