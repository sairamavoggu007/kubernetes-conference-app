package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Kubernetes conference"
	const conferenceTickets = 50
	//  unit is unsgined integer we can initialize only 0 or positive numeric value
	var remainingTickets uint = 50

	//     Array: Fixed size, how many elements the array can hold
	//     var bookings [50]string

	/*   slice is an abstraction of an array
	variable length or get an sub-array of its own
	slices are also index-based and have a size, but is resized when needed
	*/
	// dynamic list
	bookings := []string{}

	getUsers(conferenceName, conferenceTickets, remainingTickets)

	//  formatted output using Printf

	//     for remainingTickets > 0 && len(bookings) < 50 {}
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//     pointer is a variable that points to the memory address of another variable
		//     Here we are using memory address of a variable userName, email and userTickets
		//     pointers are also called special variable in Go

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter lastname name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter your tickets")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//         isValidCity := city == "San Francisco" || city == "San Jose"

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			//     bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			//     fmt.Printf("The first value is %v.\n",bookings[0])
			//     fmt.Printf("Array length %v.\n", len(bookings))
			//     fmt.Printf("Slice length %v.\n", len(bookings))

			fmt.Printf("Thankyou %v for booking %v tickets to attend %v, you will be receiving email confirmation at %v with the passes .\n", firstName, userTickets, conferenceName, email)
			fmt.Printf("Total remainingTickets available in %v are %v.\n", conferenceName, remainingTickets)

			firstNames := getFirstNames(bookings)
			fmt.Printf("FirstName of the bookings are %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our %v tickets are sold out.\n", conferenceName)
				break
			}

		} else {
			if !isValidName {
				fmt.Printf("First name %v %v you entered is too short.\n", firstName, lastName)
			}
			if !isValidEmail {
				fmt.Printf("Email %v you entered is not valid or doesn't contain @ sign.\n", email)
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets %v you entered is not valid.\n", userTickets)
			}
			//                    fmt.Printf("we only have %v, you can't book %v tickets.\n", remainingTickets, userTickets)
			//                    fmt.Printf("User Input is invalid, please try again.\n")
		}
	}
}

func getUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and  %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend \n")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	// index - to ignore a variable you don't want to use
	// iterating over the dynamic list - to get only first names
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		//             var firstName = names[0]
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

//     city := "London"
//
//     switch city{
//         case "New York":
//
//         case "Singapore":
//
//         case "San Fransicso":
//
//         case "Hyderabad", "Vijaywada":
//
//         default:
//             fmt.Println("Enter a valid city")
//
//
//     }
