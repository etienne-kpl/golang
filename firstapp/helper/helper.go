package helper

import "strings"

// we need to give the type of the parameters
// This function will return values, so we must specify their types between the () and the {}
// the capital at the begining of the function means we will epxort it
func ValidateUserInputs(firstName string, lastName string, email string, tickets uint, remainingTickets uint) (bool, bool, bool) {
	// store boolean in variable
	var isValidName bool = len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := tickets > 0 && tickets <= remainingTickets
	// Possible to return multiple values
	return isValidName, isValidEmail, isValidTickets
}
