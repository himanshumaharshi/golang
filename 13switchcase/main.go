package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Go SwitchCase")
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("You can move to spot 1")
	case 2:
		fmt.Println("You can move to spot 2")
	case 3:
		fmt.Println("You can move to spot 3")
		fallthrough // after receiving value of 3 then programme will execute next preceding case also
	case 4:
		fmt.Println("You can move to spot 4")
	case 5:
		fmt.Println("You can move to spot 5")
	case 6:
		fmt.Println("You can move to spot 6 & roll the dice again")
	default:
		fmt.Println("Try again")
	}
}
