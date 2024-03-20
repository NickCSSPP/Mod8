package main

import (
	"fmt"
	"strconv"
)

func main() {
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Steelers Conference"
	bookings := []string{}

	fmt.Println("Welcome to " + conferenceName + " booking application.\nWe have a total of " + strconv.Itoa(conferenceTickets) + " tickets available.\nGet your tickets here to attend\n")

	// Declare data types
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Collect data
	fmt.Println("Enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scanln(&userTickets)

	// Logic for booking system
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	// Output
	fmt.Printf("Thanks %v %v for booking %v tickets. You will receive a confimration email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all of our bookings: %v\n", bookings)
}
