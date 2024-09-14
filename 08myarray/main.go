package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	// declaration
	var colors [4]string

	// initialization
	colors[0] = "red"
	colors[1] = "green"
	colors[3] = "blue"

	fmt.Println("Colors: ", colors)
	fmt.Println("Colors: ", len(colors))

	// declaration + initialization
	var cars = [6]string{
		"BMW", "Mercedes", "MG", "Bugatti", "Audi", "Kia",
	}
	fmt.Println("Cars: ", cars)
	fmt.Println("Cars: ", len(cars))
}
