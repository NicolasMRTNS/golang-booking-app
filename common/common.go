package common

import "strings"

func ValidateUserInput(firstname string, lastname string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isNameValid bool = len(firstname) >= 2 && len(lastname) >= 2
	var isEmailValid bool = strings.Contains(email, "@") && strings.Contains(email, ".")
	var isTicketNumberValid bool = userTickets > 0 && userTickets <= remainingTickets
	return isNameValid, isEmailValid, isTicketNumberValid
}