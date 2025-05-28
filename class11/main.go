package main
import "fmt"
func add(a int, b int)int {
	sum:= a + b // Return the sum of a and b
	return sum
}
func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2 // Calculate the sum of num1 and num2
	multiply := num1 * num2 // Calculate the product of num1 and num2
	return sum, multiply // Return both the sum and the product	
}
func main() {
a:= 10
b:= 20	
result := add(a, b) // Call the add function with a and b
fmt.Println("The sum of", a, "and", b, "is", result) // Print the sum of a and b

sum, multiply := getNumbers(50, 60) // Call the getNumbers function with different values
fmt.Println("The sum is:", sum) // Print the sum
fmt.Println("The product is:", multiply) // Print the product
}