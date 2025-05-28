package main

import "fmt"

func checkEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is Even")
	} else {
		fmt.Println(num, "is Odd")
	}
}

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number) 

	checkEvenOdd(number) 
}
