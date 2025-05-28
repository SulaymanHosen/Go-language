package main

import "fmt"

func main() {
	age := 25
	salary := 35000

	if age >= 18 {
		if salary >= 30000 {
			fmt.Println("You are eligible for the loan")
		} else {
			fmt.Println("You need a higher salary")
		}
	} else {
		fmt.Println("You are not old enough")
	}
}
