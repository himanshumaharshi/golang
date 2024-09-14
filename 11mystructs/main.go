package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to Go Structs")
	// in golang there is no concept of inheritance and super or patent

	// use struct
	rahul := User{"Rahul", "rahul@gmail.com", true, 30}
	fmt.Println(rahul)                       // {Rahul rahul@gmail.com true 30}
	fmt.Printf("User details: %+v\n", rahul) // {Name:Rahul Email:rahul@gmail.com Status:true Age:30}
	fmt.Printf("Name: %v, Email: %v, Status: %v, Age: %v\n", rahul.Name, rahul.Email, rahul.Status, rahul.Age)
}
