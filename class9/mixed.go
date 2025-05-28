package main

import "fmt"

func main() {
	age := 22
	country := "Bangladesh"

	if (age >= 18 && country == "Bangladesh") || age >= 60 {
		fmt.Println("You are allowed to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}
}
