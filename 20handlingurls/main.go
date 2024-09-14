package main

import (
	"fmt"
	"net/url"
)

const URL = "https://dummy.restapiexample.com:3000/api/v1/employees?course=dsa&courseId=32&paymentId=29ds9jf"

func main() {
	fmt.Println("Handling URL's")
	fmt.Println("URL: ", URL)

	// Parsing URL
	result, err := url.Parse(URL)
	checkNilError(err)
	fmt.Printf("Type of parsed url: %T\n", result) // *url.URL

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("RawQuery: ", result.RawQuery) // parameters in the url (something.com/courseID=aje8f&PaymentID=dsa87)

	queryParams := result.Query()
	fmt.Println("Query: ", queryParams)
	fmt.Printf("Type of Query params: %T\n", queryParams) // url.Values (kind of maps) (in key:value format)

	//access parameters
	fmt.Println(queryParams["course"])
	for idx, param := range queryParams {
		fmt.Printf("%v:%v\n", idx, param)
	}

	// generating url from parameters
	partsOfURL := &url.URL{ // always remember to pass reference not copy
		Scheme:  "https",
		Host:    "dummy.restapiexample.com",
		Path:    "/api/v1/employees",
		RawPath: "user=himanshu",
	}
	constructedURL := partsOfURL.String()
	fmt.Println(constructedURL)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
