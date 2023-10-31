package main 

import ( "fmt"; "os"; tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5")

func main() {
	fmt.Println("Welcome")
	/* const StartTickets = 50
	var remainingTickets = 50c nf
    
	fmt.Println("Its  goapp project")
	fmt.Printf("We have %v left \n", StartTickets)
	fmt.Printf("Get your %v ", remainingTickets)

	var name string
	var tickets int
	
	name := "Tom"
	tickets := 2 */
	var name string
	var tickets int

    fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	
	fmt.Println("Enter your tickets: ")
	fmt.Scan(&tickets)

	fmt.Printf("Hello %v, get your %v tickets", name, tickets)
}
