package main

import "fmt"

func sum(num1 int, num2 int)  {
	result := num1 + num2
	fmt.Println("The sum is:", result)

}

func main() {
	a:= 10
	b:= 20
	sum(a,b)
	sum(30, 40) // Call the sum function with different values
	sum(50, 60) // Call the sum function with different values
	sum(70, 80) // Call the sum function with different values
	sum(90, 100) // Call the sum function with different values
}
