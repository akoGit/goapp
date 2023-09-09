package main

import (
	"fmt"
	"sync"
	"time"
)


const conference_tickets int = 50

var conferenceName = "Go conference"
var remaining_tickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
    firstName string
    lastName string
    email string
    numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

		    bookTicket(userTickets, firstName, lastName, email)
           
            wg.Add(1)
           
            go sendTicket(userTickets, firstName, lastName, email)
           
            firstNames := getFirstNames()
			fmt.Printf("the first names of bookings are %v\n", firstNames)

			if remaining_tickets == 0 {
				fmt.Println("SOLD OUT")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email adress should contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Printf("\x1b[31m%s\x1b[0m" , "number of tickets you entered is invalid")
            }

		}

	}
    wg.Wait()
}

func greetUsers() {
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conference_tickets, remaining_tickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remaining_tickets = remaining_tickets - userTickets

	var userData = userData {
        firstName: firstName,
        lastName: lastName,
        email: email,
        numberOfTickets: userTickets,
    }

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remaining_tickets, conferenceName)

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("enter your email adress:")
	fmt.Scan(&email)

	fmt.Println("enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
    time.Sleep(10 * time.Second)
    var ticket = fmt.Sprintf("%v ticket for %v %v", userTickets, firstName, lastName)
    fmt.Println("####################")
    fmt.Printf("Sending ticket:\n %v \nto email adress %v\n", ticket, email)
    fmt.Println("####################")
    wg.Done()
}
