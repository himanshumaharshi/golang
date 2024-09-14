package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"courseprice"`
	Category string   `json:"coursecategory"`
	Password string   `json:"-"`
	Tags     []string `json:"coursetags,omitempty"` // omitempty -> if value is nil then don't show
}

func main() {
	fmt.Println("Welcome to Golang JSON Handling")
	// encoding of JSON
	//EncodeJSON()

	// decoding of JSON
	DecodeJSON()
}

func EncodeJSON() {
	availableCourses := []course{
		{"ReactJS", 299, "video course", "asd123", []string{"web-dev", "js", "front-end"}},
		{"NodeJS", 499, "video course", "jkl123", []string{"web-dev", "js", "back-end"}},
		{"JWT", 99, "article", "uio123", nil},
	}

	// package this data as JSON data
	//finalJSON, err := json.Marshal(availableCourses)
	finalJSON, err := json.MarshalIndent(availableCourses, "", "\t") // better MarshalIndent(data, prefix, indent)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated JOSN: %s\n", finalJSON)
}

func DecodeJSON() {
	webJSONData := []byte(`
		{
                "coursename": "ReactJS",
                "courseprice": 299,
                "coursecategory": "video course",
                "coursetags": ["web-dev","js","front-end"]
        }
	`)

	var validateCourseJSON course

	checkValid := json.Valid(webJSONData)

	if checkValid {
		fmt.Println("Valid JSON")
		json.Unmarshal(webJSONData, &validateCourseJSON)
		fmt.Printf("Valid JSON: %#v\n", validateCourseJSON)
	} else {
		fmt.Println("Invalid JSON")
	}

	// add data to json key:value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(webJSONData, &myOnlineData)
	fmt.Printf("Online Data: %#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("%s: %#v\n", key, value)
	}
}
