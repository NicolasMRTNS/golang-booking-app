package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings) < 50 {
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

		// Ask the user for their email address
		fmt.Println("Please, enter your email address:")
		fmt.Scan(&email)

		// Ask the user how many tickets they would like to purchase
		fmt.Println("Please, enter the number of tickets you wish to book:")
		fmt.Scan(&userTickets)

		var isNameValid bool = len(firstname) >= 2 && len(lastname) >= 2
		var isEmailValid bool = strings.Contains(email, "@") && strings.Contains(email, ".")
		var isTicketNumberValid bool = userTickets > 0 && userTickets <= remainingTickets

		if isNameValid && isEmailValid && isTicketNumberValid {
			remainingTickets -= userTickets
			bookings = append(bookings, firstname + " " + lastname)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a 	confirmation email at %v\n", firstname, lastname, userTickets, email)
			fmt.Printf("We have %v tickets remaining for %v\n", remainingTickets, 	conferenceName)

			firstnames := []string{}
			for _, booking := range bookings {
				var names []string = strings.Fields(booking)
				firstnames = append(firstnames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstnames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, we are sold out.")
				break
			}
		} else {
			if !isNameValid {
				fmt.Println("Please, enter a valid name")
			}
			if !isEmailValid {
				fmt.Println("Please, enter a valid email address")
			}
			if !isTicketNumberValid {
				fmt.Println("Please, enter a valid number of tickets")
			}
		}
	}
}