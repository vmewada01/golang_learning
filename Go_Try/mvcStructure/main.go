package main

import (
	"fmt"
	"log"
	"mvcStructure/controllers"
	"mvcStructure/initializers"
	"net/http"

	"github.com/gorilla/mux"
)

func init(){

	initializers.ConnectToDatabase()
}

func main() {
	fmt.Println("hello world")
	fmt.Println("Server started on port 8080")

	router := mux.NewRouter()

	// Define your routes and corresponding controller functions
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	
	// Start the HTTP server with the router
	log.Fatal(http.ListenAndServe(":8080", router))
}