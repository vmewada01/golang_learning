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

type Course struct{
CourseId string `json:"courseid"`
CourseName string `json:"coursename"`
CoursePrice int `json:"price"`
Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`

}
//fake DB
var courses []Course

//middlware, helper -file

func (c *Course) IsEmpty()bool{
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseId == "" 
}

func serverHome(w http.ResponseWriter, r * http.Request){
	w.Write([]byte("<h1>Welcome to api from golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r * http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r * http.Request){
	fmt.Println("Get one course")
    w.Header().Set("Content-Type", "application/json")
	// grad id from request 

	params := mux.Vars(r)
	fmt.Println(params)
   
	//loop through courses ,find matching and return the response

	for _, course := range courses{
		if course.CourseId == params["id"]{
            json.NewEncoder(w).Encode(course)
            return
        }
	}

	json.NewEncoder(w).Encode("No course found with the given id")
    return 
}


func createOneCourse(w http.ResponseWriter, r * http.Request){
	fmt.Println("Create one course")
    w.Header().Set("Content-Type", "application/json")
   // what if  : body is empty
   if r.Body == nil {
	    json.NewEncoder(w).Encode("Please send some data")
   }
   // what about - {}

   var course Course 
   _ =  json.NewDecoder(r.Body).Decode(&course)

   if course.IsEmpty(){
        json.NewEncoder(w).Encode("No data inside JSON")
		return 
   }

   // generate a unique id, string 
   // append course into courses

   rand.Seed(time.Now().UnixNano())
   course.CourseId = strconv.Itoa(rand.Intn(100))
   courses = append(courses, course)
   json.NewEncoder(w).Encode(course)
   return 

   
}


func updateOneCourse(w http.ResponseWriter, r * http.Request){

	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// first -grab id from req

	params := mux.Vars(r)

	for index, course := range courses{
		if course.CourseId == params["id"]{

        courses = append(courses[:index], courses[index+1:]...)
    
		var course  Course

		_ = json.NewDecoder(r.Body).Decode(&course)
    
        course.CourseId= params["id"]
	    courses = append(courses, course)
		json.NewEncoder(w).Encode(course)
		return 
        }
	}



}


func deleteOneCourse(w http.ResponseWriter, r * http.Request){
   fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

  params := mux.Vars(r)

  for index, course := range courses{
    if course.CourseId == params["id"]{
      courses = append(courses[:index], courses[index+1:]...)
      json.NewEncoder(w).Encode(course)
      return 
    }
  }


}

func main() {
	fmt.Println("api version")
	r := mux.NewRouter()

	courses  = append(courses, Course{CourseId: "2", CourseName: "react", CoursePrice: 299, Author: &Author{
		Fullname: "<NAME>1",
        Website: "www.google.com",
	}})

	courses  = append(courses, Course{CourseId: "3", CourseName: "angular", CoursePrice: 399, Author: &Author{
        Fullname: "<NAME>2",
        Website: "www.google.com",
    }})

     // routing 
	 r.HandleFunc("/", serverHome).Methods("GET")
	 r.HandleFunc("/courses", getAllCourses).Methods("GET")
	 r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	 r.HandleFunc("/courses", createOneCourse).Methods("POST")
	 r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	 r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port 

	log.Fatal(http.ListenAndServe(":4000", r))
	
}