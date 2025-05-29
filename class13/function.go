// package main
// import "fmt"
// func main(){
// 	var name string
// 	fmt.Println(("Enter your name: "))
// 	fmt.Scanln(&name) // Read input from the user
// 	var num1 int
// 	var num2 int
// 	fmt.Println("Enter Your First Number: ")
// 	fmt.Scanln(&num1) // Read first number from the user	
// 	fmt.Println("Enter Your Second Number: ")
// 	fmt.Scanln(&num2) // Read second number from the user
// 	sum:= num1 + num2 // Calculate the sum of the two numbers
// 	fmt.Println("Hello", name, "Welcome to the application") // Print a welcome message
// 	fmt.Println(("the sum of the two numbers is: "), sum) // Print the sum
// fmt.Println("Thank you for using the application") // Print a thank you message
// fmt.Println("Have a great day!") // Print a closing message
// fmt.Println("Goodbye!") // Print a goodbye message
// }

package main
import "fmt"
func welcomeMessage() {
	fmt.Println( "Welcome to the application") // Print a welcome message
}
func getuserName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name) // Read input from the user
	return name
}
func getTwoNumbers() (int,int){
	var num1 int
	var num2 int
	fmt.Println("Enter Your First Number: ")
	fmt.Scanln(&num1) // Read first number from the user	
	fmt.Println("Enter Your Second Number: ")
	fmt.Scanln(&num2) // Read second number from the user
	return num1, num2
}
func addNumbers(num1 int, num2 int) int {
	sum := num1 + num2 // Calculate the sum of the two numbers
	return sum
}
func display(name string,sum int){
	fmt.Println("Hello", name, "Welcome to the application") // Print a welcome message
	fmt.Println("The sum of the two numbers is: ", sum) // Print the sum
}
func thankYouMessage() {
	fmt.Println("Thank you for using the application") // Print a thank you message
	fmt.Println("Have a great day!") // Print a closing message
	fmt.Println("Goodbye!") // Print a goodbye message
}
func main(){
welcomeMessage()
name:=getuserName()
num1,num2:=getTwoNumbers()
sum:=addNumbers(num1,num2)
display(name,sum)
thankYouMessage()
}