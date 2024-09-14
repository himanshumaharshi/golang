package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")
	presentTime := time.Now()
	fmt.Println("Current Time: ", presentTime)
	// formatting
	fmt.Println("Formated time", presentTime.Format("01-02-2006 Monday 15:04:05"))

	createDate := time.Date(2002, time.November, 12, 22, 24, 35, 01, time.UTC)
	fmt.Println("Create Date: ", createDate)

	currentNanoSec := time.Now().Nanosecond()
	fmt.Println("Current NanoSec: ", currentNanoSec)
}
