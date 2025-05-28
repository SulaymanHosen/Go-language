package main
import "fmt"
func main() {
	marks:= 85
	income:=20000
	if marks >=80 && income >= 2000 {
		fmt.Println("You are eligible for the scholarship") // Print if eligible for the scholarship
	}else if marks >= 60 && income >= 20000 {
		fmt.Println("You are eligible for the scholarship with conditions") // Print if eligible with conditions
	} else {
		fmt.Println("You are not eligible for the scholarship") // Print if not eligible for the scholarship
	}
}