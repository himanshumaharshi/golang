package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maths command in GoLang")

	//var (
	//	numberOne int     = 3
	//	numberTwo float64 = 5.8
	//)
	// fmt.Println("Sum :", numberOne + numberTwo) // Invalid operation: numberOne + numberTwo (mismatched types int and float64)
	//fmt.Println("Sum :", numberOne+int(numberTwo)) // here numberTwo will be converted to int, but we will lose precision

	// generate random number
	// 1. math/rand
	//rand.Seed(time.Now().UnixNano()) // deprecated
	//fmt.Println(rand.Intn(5))

	// 2. crypto/rand
	//randomNumber, _ := rand.Int(rand.Reader, big.NewInt(7)) // rand.Int()
	//fmt.Println(randomNumber)

	//randomNumber, err := rand.Prime(rand.Reader, 2) // rand.Prime()
	//if err != nil {
	//	fmt.Println("error: ", err)
	//} else {
	//	fmt.Println(randomNumber)
	//}

	//c := 6
	//b := make([]byte, c)
	//randomNumber, err := rand.Read(b)
	//if err != nil {
	//	fmt.Println("error: ", err)
	//} else {
	//	fmt.Println(randomNumber)
	//}
}
