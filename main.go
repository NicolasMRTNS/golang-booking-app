package main

import (
	"booking-app/common"
	"fmt"
	"sync"
	"time"
)

const conferenceName string = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings []UserData = make([]UserData, 0)

type UserData struct {
	firstname string
	lastname string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
		firstname, lastname, email, userTickets := getUserInput()

		isNameValid, isEmailValid, isTicketNumberValid := common.ValidateUserInput(firstname, lastname, email, userTickets, remainingTickets)

		if isNameValid && isEmailValid && isTicketNumberValid {
			bookTicket(userTickets, firstname, lastname, email)

			wg.Add(1)
			go sendTicket(userTickets, firstname, lastname, email)

			var firstnames []string = getFirstnames()
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v bboking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstnames() []string {
	firstnames := []string{}
	for _, booking := range bookings {
		firstnames = append(firstnames, booking.firstname)
	}
	return firstnames
}

func getUserInput() (string, string, string, uint) {
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

	return firstname, lastname, email, userTickets
}

func bookTicket(userTickets uint, firstname string, lastname string, email string) {
	remainingTickets -= userTickets

	var userData UserData = UserData {
		firstname: firstname,
		lastname: lastname,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive aconfirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Printf("We have %v tickets remaining for %v\n", remainingTickets,conferenceName)
}

func sendTicket(userTickets uint, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	var ticket string = fmt.Sprintf("%v tickets for %v %v", userTickets, firstname, lastname)
	fmt.Println("#######")
	fmt.Printf("Sending ticket:\n%v to email address: %v\n", ticket, email)
	fmt.Println("#######")
	wg.Done()
}