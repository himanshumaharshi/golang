package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go functions")
	greet()

	// normal function to add two numbers
	sum := adder(3, 5)
	fmt.Println(sum)

	// pro function to add two numbers
	total := proAdder(5, 6, 5, 2, 6, 2, 6, 2, 4, 7, 4)
	fmt.Println(total)

	totalTwo, myMessage := proAdderTwo(4, 5, 3, 6, 3, 7, 2, 6)
	fmt.Println("total: ", totalTwo)
	fmt.Println("myMessage: ", myMessage)
}

// pro functions
func proAdder(values ...int) int { // here ... are variadic functions
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func proAdderTwo(values ...int) (int, string) { // in case we want to return multiple types of values
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum, "This is Message"
}

func adder(x int, y int) int {
	return x + y
}

func greet() {
	fmt.Println("Hello, World!")
}
