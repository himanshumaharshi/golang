package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go IfElse")

	loginCount := 100
	var result string

	if loginCount > 10 {
		result = "Regular User"
	} else if loginCount < 10 {
		result = "New User"
	} else {
		result = "User Admin"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 19; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

	//if err != nil {}
}
