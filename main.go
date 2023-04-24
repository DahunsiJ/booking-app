package main

import (
	"fmt"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "GO Conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main() {

	//  Greetings to user
	greetUsers()

	// for {
		// Get user input
		firstName, lastName, email, userTickets := getUserInput()

		// Validate User Input
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// Function for remainingTickets
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			// Print First name Function
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				// Break out of the infinite loop
				fmt.Println("Our conference is booked out. Kindly Make yourself available earlier next")
				// break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email does not contain the @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Println("invalid number of ticket")
			}
		}

	// }

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets but %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// asking user for their details
	fmt.Println("Enter your firstname: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// create a struct instead of a map for a user (Key-Value pair)
	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank %v %v for booking %v tickets. You will receive a comfirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string,email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket: \n %v \n to email address %v \n", ticket, email)
	fmt.Println("####################")
}