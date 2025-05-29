package main

import "fmt"

var globalVar = 100 // Global scope

func showGlobal() {
	fmt.Println("Global variable:", globalVar)
}

func demoLocalScope() {
	localVar := 50
	fmt.Println("Local variable:", localVar)
}

func main() {
	showGlobal()     // Accessing global variable
	demoLocalScope() // Accessing local variable
	
}