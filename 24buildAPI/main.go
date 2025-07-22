package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// model for courses  - file
//

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB

var courser []Course

// middlewares -- separate file

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// controllers - differnt files 

// serve home route '/'

func serveHome(w  http.ResponseWriter , r *http.Request ) {

	w.Write([]byte("<h1> Welcome to api by Dhairya. </h1>"))
}



func getAllCourses(w http.ResponseWriter , r *http.Request) {
	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "applicatoin/json")
	json.NewEncoder(w).Encode(courser)

}







