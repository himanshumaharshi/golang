package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name: ")

	// comma ok syntax || err
	input, _ := reader.ReadString('\n')
	// input, err := reader.ReadString('\n')
	fmt.Println("Received input -> ", input)
	fmt.Printf("Type of input is %T\n", input)
}
