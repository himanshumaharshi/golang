package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go loops")
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	// normal for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// using range
	for j := range days {
		fmt.Println(days[j])
	}

	// for each loop
	for k, day := range days {
		fmt.Printf("Index is %v, and day is %v\n", k, day)
	}

	value := 1
	for value < 10 {
		if value == 6 {
			goto jump
		} else if value == 5 {
			value++
			continue
		} else if value == 7 {
			break
		} else {
			fmt.Println(value)
		}
		value++
	}

jump:
	fmt.Println("Utilized GOTO")
}
