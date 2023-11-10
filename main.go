package main

import "fmt"

func main() {
	titleName := "Meeting"
	const meetingTickets int = 100
	var remainingMeetingTickets uint = 100
	fmt.Printf("Title: %v", titleName)
	fmt.Printf("We had %v tickets, and by now there are %v avalaible",
		meetingTickets, remainingMeetingTickets)
	var varGo = "var"
	var ticketList [100] string
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of ticket that you want: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Array type: %T\n", ticketList)
	fmt.Printf("Array type: %v\n", len(ticketList))

	ticketList[0] = firstName+ " " +lastName
	remainingMeetingTickets= remainingMeetingTickets-userTickets

	fmt.Printf("Remaining tickets: %v", remainingMeetingTickets)


	fmt.Println(varGo)
}
