package main 

import ( 
	"fmt" 
    "strings"
)

func main() {
	confName := "Go conf"
	const confTickets int = 50
	remainingTickets := 50
	/* const StartTickets = 50
	var remainingTickets = 50c nf
    
	fmt.Println("Its  goapp project")
	fmt.Printf("We have %v left \n", StartTickets)
	fmt.Printf("Get your %v ", remainingTickets)

	var name string
	var tickets int
	
	name := "Tom"
	tickets := 2 
	var name string
	var tickets int

    fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	
	fmt.Println("Enter your tickets: ")
	fmt.Scan(&tickets)
*/

    fmt.Printf("Welcome to %v start app\n",confName)
	fmt.Printf("We have total of %v and %v are still available\n", confTickets, remainingTickets,)
	fmt.Printf("Hello, get your %v tickets\n", confTickets)
	
    // var booking [50]string  array of booking, because it has size defined 
    // var booking []string default define a slice called booking
    booking := []string{} // faster slice definition

	for {
		
	var name string
	var surname string
	var email string 
	var userTickets int
	
  	fmt.Println("Enter your username: ")
	fmt.Scanln(&name)

	fmt.Println("Enter your surrname: ")
	fmt.Scanln(&surname)

	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets : ")
	fmt.Scanln(&userTickets)	
	
   remainingTickets = remainingTickets - userTickets

   booking = append(booking, name + " " + surname)

   fmt.Printf("whole array %v \n", booking)
   
   fmt.Printf("Your have %v tickets left in %v\n", remainingTickets, confName)

   firstNames := []string{}
   for _, books := range booking {
	  var names = strings.Fields(books)
	  firstNames = append(firstNames, names[0])
   }
   fmt.Printf("These first names of  bookings: %v\n", firstNames)
	}

}
