package main

import "fmt"

func main() {
	age := 18
	experience := 1

	if age >= 18 || experience >= 2 {
		fmt.Println("You can apply for the job")
	} else {
		fmt.Println("You are not eligible")
	}
}
