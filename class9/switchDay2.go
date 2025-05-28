package main

import "fmt"

func main() {
	day := "Saturday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Weekend is coming")
	case "Saturday", "Sunday":
		fmt.Println("It's weekend!")
	default:
		fmt.Println("It's a normal weekday")
	}
}
