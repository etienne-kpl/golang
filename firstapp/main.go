package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	var userName string
	fmt.Println(quote.Go())
	// ask the user for their name
	// pointer & to store the variable
	fmt.Print("Please enter your name:", " ")
	fmt.Scan(&userName)
	fmt.Printf("Welcome, %v.\n", userName)
}
