package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to type conversions.")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter rating: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered rating: ", input)
	// in this we have taken the input as string and then trying to add int in it.
	// strings.TrimSpace() will remove space which was stored while getting user input
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to your rating", numRating+1)
	}
}
