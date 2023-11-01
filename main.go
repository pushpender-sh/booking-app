package main

import "fmt"

func main() {
	confName := "Go Conference"
	const confTickets = 50
	var remTickets uint = 50

	fmt.Printf("confname is of %T type  confTickets is of %T type. \n", confName, confTickets)

	fmt.Printf("Welcome to the dashboard of the %v ticket booking platform.\n", confName)
	fmt.Printf("We have a total of %v and %v are still remaining.\n", confTickets, remTickets)
	fmt.Println("Get your tickets here at best prices to attend the conference")

	var userName string
	var userTickets int

	//asked to used to give the input
	fmt.Scan((userName))
	fmt.Println(confName)
	fmt.Println(&confName)

	// userName = "Aman"
	// userTickets = 2

	fmt.Printf("User named %v booked %v tickets. \n", userName, userTickets)

}
