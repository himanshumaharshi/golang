package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to my slices")

	// slice creation
	var colors = []string{"red", "blue", "green", "orange", "black", "pink"}
	fmt.Printf("Type of colors: %T\n", colors)
	fmt.Println("Value of colors: ", colors)
	fmt.Println("Size of colors: ", len(colors))

	// adding elements into slice
	colors = append(colors, "white", "grey") // values will be added ate the end of slices
	fmt.Println("Value of colors after append: ", colors)

	// slice manipulation (targeted access of elements)
	//colors = append(colors[1:]) // starts accessing value from given number which is 1 (here 0 will be skipped) [1, 2, 3, 4, 5, 6, 7]
	//fmt.Println("Value of colors: ", colors)

	//colors = append(colors[1:3]) // starts accessing value from given index which is 1 and stops before index 3 [1, 2]
	//fmt.Println("Value of colors: ", colors)

	//colors = append(colors[:3]) // starts accessing value from index 0 and stops before 3 [0, 1, 2]
	//fmt.Println("Value of colors: ", colors)

	colors = append(colors[1:3]) // starts accessing value from given index which is 1 and stops before index 3
	fmt.Println("Value of colors: ", colors)

	// different syntax (using make(type, size))
	highScores := make([]int, 4)
	highScores[0] = 34
	highScores[1] = 43
	highScores[2] = 23
	highScores[3] = 54
	//highScores[4] = 32 // will give error

	highScores = append(highScores, 42, 43, 55) // values will be added
	fmt.Println("Value of highScores: ", highScores)

	sort.Ints(highScores) // ascending order sort
	fmt.Println("Sorted value of highScores: ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // gives true or false

	// remove a value from slice according to index
	courses := []string{"SDE", "AI", "ML", "WEB", "Crypto", "Linux"}
	fmt.Println("Courses before deletion:", courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after deletion:", courses)
}
