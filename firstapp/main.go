package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Golang lecture"
	const roomCapacity int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("The room can welcome %v guests and we have %v tickets available.\n", roomCapacity, remainingTickets)

	var firstName string
	var lastName string
	// uint: positive integer
	var tickets uint
	var bookings []string // Slice cause there is no size in the brackets

	for {
		// ask the user for their name
		// pointer & to store the variable
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("How many tickets do you want?")
		fmt.Scan(&tickets)

		remainingTickets -= tickets

		fmt.Printf("Thank you for your booking, %v %v!\n", firstName, lastName)
		fmt.Printf("You booked %v tickets.\n", tickets)
		fmt.Printf("Only %v tickets left for %v!\n", remainingTickets, conferenceName)

		bookings = append(bookings, firstName+" "+lastName)

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
}
