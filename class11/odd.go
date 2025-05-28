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

	checkEvenOdd(7)
	checkEvenOdd(12)
}
