package main

import "fmt"

func main(){
	eventName := "Tayangan Perdana Fast Furios 18: The Family" // eventName := is equal to var eventName string =
	const numberOfTickets int = 50 // fixed number
	var remainingTickets uint = 50 // remaining ticket can't be negative
	userBookings := []string{} //declare slice

	// fmt.Printf("the data type for eventName is %T, for numberOfTickets is %T, for remainingTickets is %T\n",eventName,numberOfTickets,remainingTickets)

	fmt.Printf("Hello, welcome to our booking application for %v\n", eventName)
	fmt.Println("Book your ticket now, don't miss this awesome movie!")

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

	remainingTickets = remainingTickets - userTickets
	userBookings = append(userBookings, userFirstName + " " + userLastName)

	// fmt.Printf("The whole slice:%v\n", userBookings)
	// fmt.Printf("The first value of the slice: %v\n", userBookings[0])
	// fmt.Printf("The type of the slice: %T\n", userBookings)
	// fmt.Printf("The length of the slice: %v\n", len(userBookings))

	fmt.Printf("Hello %v %v and welcome to our booking application!\n", userFirstName,userLastName)
	fmt.Printf("You (%v) just bought %v tickets.\n", userEmail, userTickets)
	fmt.Printf("We have only %v tickets and the the remaining number of tickets is %v\n",numberOfTickets, remainingTickets)

	fmt.Printf("This is the bookings in our application so far: %v\n",userBookings)
}