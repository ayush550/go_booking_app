package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Intro to Programming Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0) // slice of struct

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidEmail && isValidName && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstName()
		fmt.Printf("The bookings are: %v \n", firstNames)

	} else {
		if !isValidName {
			fmt.Printf("You have entered invalid first or last Name, make sure your name is greater than 2 characters long \n")
		}
		if !isValidEmail {
			fmt.Printf("You have entered invalid email address. Make sure it contains @ symbol \n")
		}
		if !isValidTicketNumber {
			fmt.Printf("You have entered invalid number of tickets. \n")
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are remaining to be booked \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets booked here to attend")
}

func getFirstName() []string {

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

	fmt.Println("Enter your first Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create map for user data
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("%v %v booked %v tickets. Your tickets will be sent to %v email \n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets left \n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("=====================")
	fmt.Printf("Sending ticket:\n %v\n to email address %v\n", ticket, email)
	fmt.Println("=====================")
	wg.Done()
}
