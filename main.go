package main

import (
	"fmt"
	"kubernetes-conference-app/helper"
	"time"
)

var conferenceName string = "Kubernetes conference"

const conferenceTickets = 50

// unit is unsgined integer we can initialize only 0 or positive numeric value
var remainingTickets uint = 50

//     Array: Fixed size, how many elements the array can hold
//     var bookings [50]string

/*   slice is an abstraction of an array
variable length or get an sub-array of its own
slices are also index-based and have a size, but is resized when needed
*/
// dynamic list
// Create empty list of maps
//var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

// Here UserData stuct is custom data type
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	getUsers()

	//  formatted output using Printf

	//     for remainingTickets > 0 && len(bookings) < 50 {}
	for {
		//     pointer is a variable that points to the memory address of another variable
		//     Here we are using memory address of a variable userName, email and userTickets
		//     pointers are also called special variable in Go

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//         isValidCity := city == "San Francisco" || city == "San Jose"

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
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

func getUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and  %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend \n")
}

func getFirstNames() []string {
	firstNames := []string{}
	// index - to ignore a variable you don't want to use
	// iterating over the dynamic list - to get only first names
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		//             var firstName = names[0]
		//firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter lastname name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter your tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	//var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//userData["firstName"] = firstName
	//userData["lastnName"] = lastName
	//userData["email"] = email
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//     bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)
	fmt.Printf("List of booking  is %v.\n", bookings)

	//     fmt.Printf("The first value is %v.\n",bookings[0])
	//     fmt.Printf("Array length %v.\n", len(bookings))
	//     fmt.Printf("Slice length %v.\n", len(bookings))

	fmt.Printf("Thankyou %v for booking %v tickets to attend %v, you will be receiving email confirmation at %v with the passes .\n", firstName, userTickets, conferenceName, email)
	fmt.Printf("Total remainingTickets available in %v are %v.\n", conferenceName, remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v to your email address %v \n", ticket, email)
	fmt.Println("################")
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
