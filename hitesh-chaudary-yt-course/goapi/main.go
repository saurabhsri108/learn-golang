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

type Course struct {
	CourseId    string  `json:"id"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// middleware/helper functions
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// controllers -> how to handle a situation GET, POST, DELETE, UPDATE

func serveHomeRoute(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to Go API Learning</h1>"))
}

func getAllCourses(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Get all the courses")
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(courses)
}

func getCourseById(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Get one course")
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	fmt.Println(params)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(res).Encode(course)
			return
		}
	}
	json.NewEncoder(res).Encode("No course found with given id")
}

func createSingleCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Create Single course")
	res.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if req.Body == nil {
		json.NewEncoder(res).Encode("No data sent from client")
		return
	}

	// what about: body is {}
	var course Course
	_ = json.NewDecoder(req.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(res).Encode("No data sent from client")
		return
	}

	// generate unique id, convert it to string
	// append course into courses
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseId = strconv.Itoa(r.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(res).Encode(course)
}

func updateSingleCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Update Single course")
	res.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if req.Body == nil {
		json.NewEncoder(res).Encode("No data sent from client")
		return
	}

	// what about: body is {}
	var course Course
	_ = json.NewDecoder(req.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(res).Encode("No data sent from client")
		return
	}

	params := mux.Vars(req)

	fmt.Println(params, course)
	// loop, id, remove, add with my id
	for index, _course := range courses {
		fmt.Println(_course.CourseId, params["id"], _course.CourseId == params["id"])
		if _course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			course.CourseId = _course.CourseId
			fmt.Println("Removal courses:", courses, course)
			courses = append(courses, course)
			fmt.Println("Appended courses:", courses)
			json.NewEncoder(res).Encode(course)
			return
		}
	}

	// send a response when id is not found
	json.NewEncoder(res).Encode("Course doesn't exist")
}

func deleteSingleCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Delete Single course")
	res.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if req.Body == nil {
		json.NewEncoder(res).Encode("No data sent from client")
		return
	}

	// what about: body is {}
	var course Course
	_ = json.NewDecoder(req.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(res).Encode("No data sent from client")
		return
	}

	params := mux.Vars(req)

	// loop, id, remove, add with my id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(res).Encode("Course deleted")
			return
		}
	}

	// send a response when id is not found
	json.NewEncoder(res).Encode("Course doesn't exist")
}

func main() {
	route := mux.NewRouter()
	// seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Vasu Srivastava", Website: "vasusrivastava.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "MERN", CoursePrice: 299, Author: &Author{Fullname: "Vasu Srivastava", Website: "vasusrivastava.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Angular", CoursePrice: 299, Author: &Author{Fullname: "Vasu Srivastava", Website: "vasusrivastava.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Golang", CoursePrice: 299, Author: &Author{Fullname: "Vasu Srivastava", Website: "vasusrivastava.com"}})

	// routing
	route.HandleFunc("/", serveHomeRoute).Methods("GET")
	route.HandleFunc("/courses", getAllCourses).Methods("GET")
	route.HandleFunc("/courses/{id}", getCourseById).Methods("GET")
	route.HandleFunc("/courses", createSingleCourse).Methods("POST")
	route.HandleFunc("/courses/{id}", updateSingleCourse).Methods("PUT")
	route.HandleFunc("/courses/{id}", deleteSingleCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", route))
}
