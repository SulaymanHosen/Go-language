package main
import "fmt"
func main(){
	fmt.Println("Welcome to the application")
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name) // Read input from the user
	fmt.Println("Hello", name, "Welcome to the application") // Print a welcome message
}		