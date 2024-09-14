package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var ptr *int
	fmt.Println("Default value of Pointer: ", ptr) // <nil>
	fmt.Printf("ptr type: %T\n", ptr)

	var one = 1
	ptr = &one
	fmt.Println("Pointed address of Pointer: ", ptr)
	fmt.Println("Pointed value of Pointer: ", *ptr)

	*ptr = *ptr + 78
	fmt.Println("updated value by Pointer: ", *ptr)
}
