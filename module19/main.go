package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("My Go API")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "3", CourseName: "GO", CoursePrice: "999", Author: &Author{Fullname: "Ayush Khare", Website: "mugenshouren.com/go"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Python", CoursePrice: "999", Author: &Author{Fullname: "Ayush Khare", Website: "mugenshouren.com/python"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get the course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create the course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	// generate a unique ID, string

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// append course into courses

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update the course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var courseRef Course
			_ = json.NewDecoder(r.Body).Decode(&courseRef)
			courseRef.CourseId = params["id"]
			courses = append(courses, courseRef)
			json.NewEncoder(w).Encode("Course Updated")
			return
		}
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted")
			return
		}
	}
}
