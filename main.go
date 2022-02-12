package main

import "fmt"

func main() {
	const conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstname string
	var lastname string
	var email string
	var userTickets uint
	// Ask the user for their first name
	fmt.Println("Please, enter your first name:")
	fmt.Scan(&firstname)

	// Ask the user for their last name
	fmt.Println("Please, enter your last name:")
	fmt.Scan(&lastname)

	// Ask the user for their email adress
	fmt.Println("Please, enter your email adress:")
	fmt.Scan(&email)

	// Ask the user how many tickets they would like to purchase
	fmt.Println("Please, enter the number of tickets you wish to book:")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Printf("We have %v tickets remaining for %v\n", remainingTickets, conferenceName)
}