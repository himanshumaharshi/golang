package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go Methods!")

	rahul := User{"Rahul", "rahul@gmail.com", true, 30}
	//fmt.Println(rahul)                       // {Rahul rahul@gmail.com true 30}
	//fmt.Printf("User details: %+v\n", rahul) // {Name:Rahul Email:rahul@gmail.com Status:true Age:30}
	fmt.Printf("Name: %v, Email: %v, Status: %v, Age: %v\n", rahul.Name, rahul.Email, rahul.Status, rahul.Age)

	// calling method
	rahul.GetStatus()
	rahul.ChangeEmail()
	fmt.Printf("Name: %v, Email: %v, Status: %v, Age: %v\n", rahul.Name, rahul.Email, rahul.Status, rahul.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// GetStatus defining a method
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) ChangeEmail() {
	u.Email = "rahul2@gmail.com"
	fmt.Println(u.Email)
}
