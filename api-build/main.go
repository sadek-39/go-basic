package main

import (
	"api-basic/router"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice float64 `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fakeDB
var courses []Course

// helper function
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func main() {

	r := router.Router()
	fmt.Println("Server is starting...")

	//r := mux.NewRouter()
	//
	////seeding
	//courses = append(courses, Course{CourseId: "2", CourseName: "React", CoursePrice: 20, Author: &Author{Fullname: "sadek", Website: "www.google.com"}})
	//courses = append(courses, Course{CourseId: "3", CourseName: "Java", CoursePrice: 20, Author: &Author{Fullname: "sadek", Website: "www.google.com"}})
	//
	////routing
	//
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", GetSingleCourse).Methods("GET")
	r.HandleFunc("/create-course", CreateCourse).Methods("POST")
	r.HandleFunc("/update-course/{id}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/delete-course/{id}", DeleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening to port 4000....")

}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api route</h1>"))
}

func GetAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Single Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request

	params := mux.Vars(r)
	fmt.Println("Course id ", params["id"])

	//loop through courses, find matching id and return the response

	for _, course := range courses {
		fmt.Println("Course id in loop and outside if ", course.CourseId)

		if course.CourseId == params["id"] {
			fmt.Println("Course id in loop and if ", params["id"])
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with this id")
	return
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course")
	w.Header().Set("Content-Type", "application/json")

	//if body empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		} else {
			json.NewEncoder(w).Encode("Course not found")
			return
		}
	}
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted success")
			break
		} else {
			json.NewEncoder(w).Encode("Course not found")
			return
		}
	}
}
