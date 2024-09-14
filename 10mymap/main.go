package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Go Maps")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["CSS"] = "CSS"
	languages["HTML"] = "HTML"
	languages["Java"] = "Java"
	languages["TS"] = "Typescript"

	fmt.Println("List of all languages", languages)
	fmt.Printf("The type of languages is %T\n", languages)

	delete(languages, "JS")
	fmt.Println("List of all languages", languages)

	for key, value := range languages {
		fmt.Println("Key:", key, "Value:", value)
	}
}
