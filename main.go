package main

import "fmt"

func main(){
	eventName := "Tayangan Perdana Fast Furios 18: The Family" // eventName := is equal to var eventName string =
	const numberOfTickets int = 50
	var remainingTickets uint = 50 // remaining ticket can't be negative

	fmt.Printf("the data type for eventName is %T, for numberOfTickets is %T, for remainingTickets is %T\n",eventName,numberOfTickets,remainingTickets)

	fmt.Printf("Hello, welcome to our booking application for %v\n", eventName)
	fmt.Printf("We have only %v tickets and the the remaining number of tickets is %v\n",numberOfTickets, remainingTickets)
	fmt.Println("Book your ticket now, don't miss this awesome movie!")

	var userName string
	var userTickets int

	userName = "Luqman"
	userTickets = 3

	fmt.Printf("%v just bought %v tickets just now\n", userName, userTickets)
}
