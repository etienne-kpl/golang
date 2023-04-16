package main

import (
	"firstapp/helper"
	"fmt"
	"strconv"
)

// Package level variables here:
// we cannot use this short writing in package level variables
// conferenceName := "Golang lecture"
// we could capitalize a package level variable to make it acessible outside the package
const conferenceName = "Golang lecture"
const roomCapacity int = 50

var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0) // alternate way of writing, 0 being the initial size of the slice
var bookings []map[string]string // Slice cause there is no size in the brackets

func main() {

	greetUsers()

	// for as long as this is "true", run the loop
	for remainingTickets > 0 && len(bookings) < 50 {

		// Store the values returned by the func
		firstName, lastName, email, tickets := getUserInput()
		// we call the function in the helper package
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInputs(firstName, lastName, email, tickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTicket(tickets, firstName, lastName, email)

			printBookings()

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

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("The room can welcome %v guests and we have %v tickets available.\n", roomCapacity, remainingTickets)
}

func printBookings() {
	fmt.Printf("There are %v bookings:\n", len(bookings))

	// range demande toujours 2 paramètres
	// si on ne se sert pas de l'index, on doit mettre un _ à la place pour le dire au compileur
	// for _, booking := range bookings {
	// }
	for index, booking := range bookings {
		// var names = strings.Fields(booking)
		// fmt.Printf("%v. %v %v.\n", index+1, names[0], names[1][0:1])
		fmt.Printf("%v. %v %v. - %v tickets\n", index+1, booking["firstName"], booking["lastName"][0:1], booking["tickets"])
	}
}

func getUserInput() (string, string, string, uint) {
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

func bookTicket(tickets uint, firstName string, lastName string, email string) {
	remainingTickets -= tickets

	// create a map for a user
	// only one data type in a map!
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	// convert the uint into a string
	userData["tickets"] = strconv.FormatUint(uint64(tickets), 10)

	fmt.Printf("Thank you for your booking, %v %v!\n", firstName, lastName)
	fmt.Printf("You booked %v tickets. You will recieve an email at %v\n", tickets, email)
	fmt.Printf("Only %v tickets left for %v!\n", remainingTickets, conferenceName)

	bookings = append(bookings, userData)
}
