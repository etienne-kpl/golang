package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Golang lecture"
	const roomCapacity int = 50
	var remainingTickets uint = 50
	var bookings []string // Slice cause there is no size in the brackets

	greetUsers(conferenceName, roomCapacity, remainingTickets)

	// for as long as this is "true", run the loop
	for remainingTickets > 0 && len(bookings) < 50 {

		// Store the values returned by the func
		firstName, lastName, email, tickets := getUserInput(remainingTickets)
		isValidName, isValidEmail, isValidTickets := validateUserInputs(firstName, lastName, email, tickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTicket(remainingTickets, tickets, firstName, lastName, email, conferenceName, bookings)

			printBookings(bookings)

			if remainingTickets == 0 {
				fmt.Printf("%v is fully booked, sorry!\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered is not valid.")
			}
			if !isValidTickets {
				fmt.Printf("Number of tickets is invalid.")
			}
			// // next in ruby: skip to the next loop
			// continue
		}
	}
}

// we need to give the type of the parameters
func greetUsers(conferenceName string, roomCapacity int, remainingTickets uint) {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("The room can welcome %v guests and we have %v tickets available.\n", roomCapacity, remainingTickets)
}

func printBookings(bookings []string) {
	fmt.Printf("There are %v bookings:\n", len(bookings))

	// range demande toujours 2 paramètres
	// si on ne se sert pas de l'index, on doit mettre un _ à la place pour le dire au compileur
	// for _, booking := range bookings {
	// }
	for index, booking := range bookings {
		var names = strings.Fields(booking)
		fmt.Printf("%v. %v %v.\n", index+1, names[0], names[1][0:1])
	}
}

// This function will return values, so we must specify their types between the () and the {}
func validateUserInputs(firstName string, lastName string, email string, tickets uint, remainingTickets uint) (bool, bool, bool) {
	// store boolean in variable
	var isValidName bool = len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := tickets > 0 && tickets <= remainingTickets
	// Possible to return multiple values
	return isValidName, isValidEmail, isValidTickets
}

func getUserInput(remainingTickets uint) (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	// uint: positive integer
	var tickets uint

	// ask the user for their name
	// pointer & to store the variable
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Printf("%v tickets left. How many do you want to book?\n", remainingTickets)
	fmt.Scan(&tickets)
	return firstName, lastName, email, tickets
}

func bookTicket(remainingTickets uint, tickets uint, firstName string, lastName string, email string, conferenceName string, bookings []string) {
	remainingTickets -= tickets

	fmt.Printf("Thank you for your booking, %v %v!\n", firstName, lastName)
	fmt.Printf("You booked %v tickets. You will recieve an email at %v\n", tickets, email)
	fmt.Printf("Only %v tickets left for %v!\n", remainingTickets, conferenceName)

	bookings = append(bookings, firstName+" "+lastName)
	return bookings
}
