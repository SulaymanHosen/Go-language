package main
import "fmt"
func main(){
	var name string
	fmt.Println(("Enter your name: "))
	fmt.Scanln(&name) // Read input from the user
	var num1 int
	var num2 int
	fmt.Println("Enter Your First Number: ")
	fmt.Scanln(&num1) // Read first number from the user	
	fmt.Println("Enter Your Second Number: ")
	fmt.Scanln(&num2) // Read second number from the user
	sum:= num1 + num2 // Calculate the sum of the two numbers
	fmt.Println("Hello", name, "Welcome to the application") // Print a welcome message
	fmt.Println(("the sum of the two numbers is: "), sum) // Print the sum
fmt.Println("Thank you for using the application") // Print a thank you message
fmt.Println("Have a great day!") // Print a closing message
fmt.Println("Goodbye!") // Print a goodbye message
}