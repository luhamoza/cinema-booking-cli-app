package main

import "fmt"

func main(){
	var eventName = "Tayangan Perdana Fast Furios 18: The Family"
	const numberOfTickets = 50
	var remainingTickets = 50

	fmt.Printf("Hello, welcome to our booking application for %v\n", eventName)
	fmt.Printf("We have only %v tickets and the the remaining number of tickets is %v\n",numberOfTickets, remainingTickets)
	fmt.Println("Book your ticket now, don't miss this awesome movie!")
}
