package main

import (
	"fmt"
	"strings"
)

const numberOfTickets int = 50                                        // fixed number
var eventName string = "Tayangan Perdana Fast Furious 18: The Family" // eventName := is equal to var eventName string =
var remainingTickets uint = 50                                        // remaining ticket can't be negative
var userBookings = []string{}                                         // declare slice

func main() {

	greetUser()

	for {
		userFirstName, userLastName, userEmail, userTickets := getUserInput()
		validUserName, validEmail, validNumberTickets := validateUserInput(userFirstName, userLastName, userEmail, userTickets)

		if validUserName && validEmail && validNumberTickets {

			bookingTicket(userTickets, userFirstName, userLastName, userEmail)

			getFirstName := getUserFirstName()
			fmt.Printf("This is the bookings in our application so far: %v\n", getFirstName)

			var lastTickets bool = remainingTickets == 0
			if lastTickets {
				fmt.Println("The tickets are sold out. Please try again next year.")
				break
			}
		} else {
			if !validUserName {
				fmt.Println("You have entered first name or last name too short")
			}
			if !validEmail {
				fmt.Println("Please enter valid email with @ sign")
			}
			if !validNumberTickets {
				fmt.Printf("Please try again with number of tickets less than %v.\n", remainingTickets)
				fmt.Printf("You cannot book %v tickets because we only have %v tickets\n", userTickets, remainingTickets)
			}
		}
	}
}

func greetUser() {
	fmt.Printf("Hello, welcome to our booking application for %v\n", eventName)
	fmt.Println("Book your ticket now, don't miss this awesome movie!")
}

func getUserInput() (string, string, string, uint) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint

	fmt.Println("Please enter your first name:")
	fmt.Scan(&userFirstName)
	fmt.Println("Please enter your second name:")
	fmt.Scan(&userLastName)
	fmt.Println("Please enter your email:")
	fmt.Scan(&userEmail)
	fmt.Println("Number of tickets you want to purchase")
	fmt.Scan(&userTickets)

	return userFirstName, userLastName, userEmail, userTickets
}

func validateUserInput(firstName string, lastName string, email string, tickets uint) (bool, bool, bool) {
	validUserName := len(firstName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@")
	validNumberTickets := remainingTickets >= tickets && tickets > 0
	return validUserName, validEmail, validNumberTickets
}

func bookingTicket(tickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - tickets
	userBookings = append(userBookings, firstName+" "+lastName+",")
	fmt.Printf("Hello %v %v and welcome to our booking application!\n", firstName, lastName)
	fmt.Printf("You (%v) just bought %v tickets.\n", email, tickets)
	fmt.Printf("We have only %v tickets and the the remaining number of tickets is %v\n", numberOfTickets, remainingTickets)
}

func getUserFirstName() []string {
	userFirstNames := []string{}
	for _, booking := range userBookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		userFirstNames = append(userFirstNames, firstName+",")
	}
	return userFirstNames
}
