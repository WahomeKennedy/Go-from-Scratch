package main // package name. Our Go application should start with package main

import "fmt"

func main() { // main function. The entry point of our Go application
	var conferenceName string = "GopherCon"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)                                      // fmt or format is a package that provides I/O functions
	fmt.Printf("We have a total of %v and %v are still available.\n", conferenceTickets, remainingTickets) //we use Printf when we want to combine the variables with the string.
	fmt.Println("Get your tickets now to attend")

	var userName string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	userTickets = 25

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
