package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Get Request in Golang")

	PerformGetRequest()      // get request
	PerformPostJSONRequest() // post request
	PerformPostFormRequest() // postForm request (encoded)
}

func PerformGetRequest() {
	const getURL = "http://localhost:3000/get"
	res, err := http.Get(getURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK { // 200
		fmt.Printf("status code error: %d %s", res.StatusCode, res.Status)
	}

	fmt.Println("StatusCode: ", res.StatusCode)
	fmt.Println("ContentLength: ", res.ContentLength)

	// method 1: convert received data bytes into string
	//dataByte, _ := io.ReadAll(res.Body)
	//data := string(dataByte)
	//fmt.Println("Data: ", data)

	// method 2: another method to use strings method
	var responseString strings.Builder
	dataByte, _ := io.ReadAll(res.Body)
	byteCount, _ := responseString.Write(dataByte)
	fmt.Println("byteCount: ", byteCount)
	fmt.Println("Response: ", responseString.String())
}

func PerformPostJSONRequest() {
	const postURL = "http://localhost:3000/post"

	// fake JSON payload
	requestBody := strings.NewReader(`
		{
			"name":"Himanshu Maharshi",
			"position":"Associate Software Development Engineer"
		}
	`)
	res, err := http.Post(postURL, "application/json", requestBody)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}

	var responseString strings.Builder
	dataBytes, _ := io.ReadAll(res.Body)
	byteCount, _ := responseString.Write(dataBytes)
	fmt.Println("byteCount: ", byteCount)
	fmt.Println("Response: ", responseString.String())
}

func PerformPostFormRequest() {
	const postFormURL = "http://localhost:3000/postform"
	data := url.Values{}
	data.Add("name", "Himanshu Maharshi")
	data.Add("position", "Associate Software Development Engineer")

	res, err := http.PostForm(postFormURL, data)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}

	var responseString strings.Builder
	dataBytes, _ := io.ReadAll(res.Body)
	byteCount, _ := responseString.Write(dataBytes)
	fmt.Println("byteCount: ", byteCount)
	fmt.Println("Response: ", responseString.String())
}
