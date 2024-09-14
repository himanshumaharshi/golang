package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Go Files")

	// content for file
	content := "Hello my name is Himanshu Maharshi"

	file, err := os.Create("./myFile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content) // gives length of file as output
	checkNilError(err)

	fmt.Println("Length is: ", length)

	// close file
	defer file.Close()

	// read file
	readFile("./myFile.txt")
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Data inside file: ", string(dataByte))
	fmt.Println("Raw data of file: ", dataByte)
}

// best practice to check errors
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
