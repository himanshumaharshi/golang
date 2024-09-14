package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://dummy.restapiexample.com/api/v1/employees"

func main() {
	fmt.Println("Welcome to Go Web Requests")

	res, err := http.Get(URL)
	checkNilError(err)

	fmt.Printf("Response is of type %T\n", res)

	defer res.Body.Close() // caller's responsibility to close the connection

	dataBytes, err := io.ReadAll(res.Body)
	checkNilError(err)
	content := string(dataBytes)
	fmt.Println("Response.Body: ", content)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
