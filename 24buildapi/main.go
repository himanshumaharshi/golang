package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Course model for courses - supposed to be in a separate file
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Name    string `json: "author_name"`
	Website string `json: "website"`
}

// fake DB
var courses []Course

// IsEmpty middleware, helper - supposed to be in a separate file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == "" && c.CoursePrice == 0
}

func main() {
	fmt.Println("APIs in Golang")
}

// ServeHome controllers - supposed to be in a separate file
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Home Page!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting course by ID")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return response
	for _, item := range courses {
		if item.CourseId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found by corresponding ID")
	return
}
