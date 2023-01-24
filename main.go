package main

import "fmt"

func _main() {
	var conferenceName string = "Go Conference"
	// shorthand only for var, cannot use this for const
	// conferenceName := "Go Conference"  //Create a variable and assign value with type
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	// array (fixed type)
	// var bookings = [50] string {}
	// var bookings [50]string
	// bookings := [50]string{}

	// no fixed size -> this is called "slices"
	// bookings := []string {}
	// var bookins = [] string {}
	var bookings []string

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	// fmt.Println("Get your tickets here to attend")

	fmt.Printf("conferenceName type is %T, conferenceTickets type is %T and remainingTickets type is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask the user for their name

	fmt.Print("Enter your first name: ")
	// & this is a pointer
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	// bookings[0] = firstName + " " + lastName

	// add a new element to bookings slice
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("The whole array %v\n", bookings)
	// fmt.Printf("The first array %v\n", bookings[0])
	// fmt.Printf("The array type %T\n", bookings)
	// fmt.Printf("Array length %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all our bookings %v\n", bookings)
}
