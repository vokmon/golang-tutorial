package main

// golang-tutorial is the name declared in go.mod
import (
	"fmt"
	"golang-tutorial/helper"
	"sync"
	"time"
)

// Package level variables, no need to pass these variables around.
// All function in the same package can access them
const conferenceTickets uint = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = conferenceTickets

// The size will be increased automatically when appending a new value to the slice
// var bookings = make([]map[string]string, 0)

// Create an empty list of struct
var bookings = make([]UserData, 0)

// struct
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// fmt.Printf("conferenceName type is %T, conferenceTickets type is %T and remainingTickets type is %T\n", conferenceName, conferenceTickets, remainingTickets)
	greetUser()

	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	// fmt.Println("Get your tickets here to attend\n")

	// for with condition + type casting
	// for remainingTickets > 0 && len(bookings) <= int(conferenceTickets) {
	// infinite loop
	for {
		// get the return value form the function
		firstName, lastName, email, userTickets := getUserInput()

		// get the return value form the function
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// if userTickets > remainingTickets {
		if !isValidName || !isValidEmail || !isValidTicketNumber {
			// fmt.Println("We only have %v tickets remaining, so you cannot book %v ticket.", remainingTickets, userTickets)
			if !isValidName {
				fmt.Println("First name or last name is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you enter is invalid.")
			}
			continue
		} else if userTickets == remainingTickets {
			// ...
		} else {
			// ...
		}

		bookTicket(userTickets, firstName, lastName, email)

		// keyword "go" make the function run in the separate thread
		// wg.Add -> number of thread to wait
		// wg.Add(2) when ther is another go sendTicket2 and to on
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		// fmt.Printf("These are all our bookings %v\n\n\n", bookings)
		fmt.Printf("The first names of the bookings are %v\n\n\n", firstNames)

		noTicketRemaining := remainingTickets == 0
		if noTicketRemaining {
			// end the program
			fmt.Println("Our conference is booked out. Comeback next year.")
			break
		}

		// Wait for all threads
		wg.Wait()
	}

	// Switch statement
	// city := "London"

	// switch city {
	// case "New York":
	// 	// Execute code for New York
	// case "Singapore", "Hong Kong":
	// 	// Execute code for Singapore
	// case "London", "Berlin":
	// 	// Execute code for London and Berlin
	// case "Mexico City":
	// 	// Execute code for Mexico City
	// default:
	// 	// Execute the default case
	// }
}

func greetUser() {
	// Use the package variables declared above

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	// for index, booking := range bookings {
	// _ ignore the variable.
	for _, booking := range bookings {

		// var names = strings.Fields(booking)
		// var firstName = names[0]
		// firstNames = append(firstNames, firstName)

		// var firstName = booking["firstName"]
		var firstName = booking.firstName
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	// & this is a pointer
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for users. Map of string with value of string
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	// format string
	var ticket = fmt.Sprintln("%v tickets for %v %v.", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Println("Sending ticket: %v \nto email address %v\n", ticket, email)
	fmt.Println("####################")
	wg.Done()
}
