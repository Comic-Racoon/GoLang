package main

import "fmt"

func main() {
	//	greet the customers and provide information about tickets available and print them in cli

	var conferenceName = "Go Lang Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to ", conferenceName, " Booking Application")
	fmt.Printf("We have total of %v  tickets  and %v tikets are still available. \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend")

}
