package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const welcomeDisplayName string = "Booking-App for conference"

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg sync.WaitGroup

func main() {

	printWelcomeMessage()

	var userData UserData
	var isValidNameLen bool = true
	var isValidEmail bool = true

	helper.CheckForLoopWithArray()

	userData.firstName, userData.lastName, userData.email, userData.numberOfTickets = scanInputValues()

	isValidNameLen = helper.ValidateNames(userData.firstName, userData.lastName)

	if !isValidNameLen {
		fmt.Println("Names are not in valid length")
		return
	}

	isValidEmail = helper.ValidateEmailID(userData.email)
	if !isValidEmail {
		fmt.Println("Email ID is not valid type ")
		return
	}

	displayUserDetails(userData.firstName, userData.lastName, userData.email, userData.numberOfTickets)

	wg.Add(3)
	go sendEmailNotification()
	go sendSMSNotification()
	go sendPushNotification()
	wg.Wait()

	helper.WorkWithMap()
}

func displayUserDetails(firstName string, lastName string, email string, numberOfTickets uint) {

	fmt.Printf("\n Fist Name is : %v", firstName)
	fmt.Printf("\n Last Name is : %v", lastName)
	fmt.Printf("\n Email ID is : %v", email)
	fmt.Printf("\n Number of tickets are  : %v\n", numberOfTickets)

}

func printWelcomeMessage() {
	fmt.Printf("Welcome to %v \n", welcomeDisplayName)
}

func scanInputValues() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var numberOfTickets uint

	fmt.Println("Enter user firstName:")
	fmt.Scan(&firstName)

	fmt.Println("Enter user lastName :")
	fmt.Scan(&lastName)

	fmt.Println("Enter user email ID :")
	fmt.Scan(&email)

	fmt.Println("Enter user number of tickets:")
	fmt.Scan(&numberOfTickets)

	return firstName, lastName, email, numberOfTickets
}

func sendEmailNotification() {
	time.Sleep(10 * time.Second)
	fmt.Println("###############################")
	fmt.Println("Email Notificaiton Sent to booking-app customer")
	fmt.Println("###############################")
	wg.Done()
}

func sendSMSNotification() {
	time.Sleep(10 * time.Second)
	fmt.Println("###############################")
	fmt.Println("SMS Notificaiton Sent to booking-app customer")
	fmt.Println("###############################")
	wg.Done()
}

func sendPushNotification() {
	time.Sleep(10 * time.Second)
	fmt.Println("###############################")
	fmt.Println("Push Notificaiton Sent to booking-app customer")
	fmt.Println("###############################")
	wg.Done()
}
