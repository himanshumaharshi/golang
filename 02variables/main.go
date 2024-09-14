package main

import "fmt"

// in GO if we use first letter as CAPITAL in variable name then it is set to be Public and can be accessed from outside also

// LoginToken constant variables
const LoginToken string = "sda4kj3hi3d9nw87dunk3w4jbw9uev" // Public

func main() {
	var userName = "Himanshu"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T\n", userName)

	var isLoggedIn = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallInt uint8 = 255 // uint -> 0 - 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type: %T\n", smallInt)

	var smallFloat float32 = 255.786487367 // float32 -> only 5 preceding values after decimal
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	var bigFloat = 255.7864873693827789 // float64 -> only 14 preceding values after decimal
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T\n", bigFloat)

	// default values and aliases
	/*
		int -> 0
		float -> 0
		string -> ""
		bool -> false
		complex -> (0+0i)
	*/
	var anotherVariable complex128
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	// implicit type
	var website = "https://rebrand.ly/himanshu-maharshi" // here lexer will automatically identify the type of variable
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T\n", website)

	// walrus operator (:=)
	rollNo := 78 // only allowed to be used inside a method only
	fmt.Println(rollNo)
	fmt.Printf("Variable is of type: %T\n", rollNo)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)
}
