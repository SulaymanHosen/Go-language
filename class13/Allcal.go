package main

import "fmt"

// Welcome Message
func welcomeMessage() {
	fmt.Println("🔷 Welcome to the Calculator Application")
}

// Get user name
func getUserName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

// Get two numbers
func getTwoNumbers() (int, int) {
	var num1, num2 int
	fmt.Print("Enter First Number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter Second Number: ")
	fmt.Scanln(&num2)
	return num1, num2
}

// Add
func add(a int, b int) int {
	return a + b
}

// Subtract
func subtract(a int, b int) int {
	return a - b
}

// Multiply
func multiply(a int, b int) int {
	return a * b
}

// Divide
func divide(a int, b int) float64 {
	if b == 0 {
		fmt.Println("⚠️ Cannot divide by zero!")
		return 0
	}
	return float64(a) / float64(b)
}

// Display all results
func displayResults(name string, a int, b int) {
	fmt.Println("\n👋 Hello", name)
	fmt.Println("📌 Here are your calculation results:")
	fmt.Println("Addition:", add(a, b))
	fmt.Println("Subtraction:", subtract(a, b))
	fmt.Println("Multiplication:", multiply(a, b))
	fmt.Println("Division:", divide(a, b))
}

// Thank you message
func thankYouMessage() {
	fmt.Println("\n🙏 Thank you for using the calculator!")
	fmt.Println("Have a nice day! 😊")
}

func main() {
	welcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	displayResults(name, num1, num2)
	thankYouMessage()
}
