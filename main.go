package main

import (
	"fmt"
	"strings"
)

func main() {
	eventName := "Tayangan Perdana Fast Furious 18: The Family" // eventName := is equal to var eventName string =
	const numberOfTickets int = 50                              // fixed number
	var remainingTickets uint = 50                              // remaining ticket can't be negative
	userBookings := []string{}                                  //declare slice

	fmt.Printf("Hello, welcome to our booking application for %v\n", eventName)
	fmt.Println("Book your ticket now, don't miss this awesome movie!")

	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint

	for {
		fmt.Println("Please enter your first name:")
		fmt.Scan(&userFirstName)
		fmt.Println("Please enter your second name:")
		fmt.Scan(&userLastName)
		fmt.Println("Please enter your email:")
		fmt.Scan(&userEmail)
		fmt.Println("Number of tickets you want to purchase")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			userBookings = append(userBookings, userFirstName+" "+userLastName+",")

			fmt.Printf("Hello %v %v and welcome to our booking application!\n", userFirstName, userLastName)
			fmt.Printf("You (%v) just bought %v tickets.\n", userEmail, userTickets)
			fmt.Printf("We have only %v tickets and the the remaining number of tickets is %v\n", numberOfTickets, remainingTickets)

			userFirstNames := []string{}
			for _, booking := range userBookings {
				var names = strings.Fields(booking)
				var firstName = names[0]
				userFirstNames = append(userFirstNames, firstName+",")
			}
			fmt.Printf("This is the bookings in our application so far: %v\n", userFirstNames)

			var lastTickets bool = remainingTickets == 0
			if lastTickets {
				fmt.Println("The tickets are sold out. Please try again next year.")
				break
			}
		} else {
			fmt.Printf("You cannot book %v tickets because we only have %v tickets\n", userTickets, remainingTickets)
			fmt.Printf("Please try again with number of tickets less than %v.\n", remainingTickets)
			continue
		}

	}
}
