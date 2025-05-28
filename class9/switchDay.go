package main

import "fmt"

func main() {
	day := 4

	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Invalid day")
	}
}
