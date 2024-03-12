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
	
	fmt.Println("Server started on port 8080")

	router := mux.NewRouter()

	// Define your routes and corresponding controller functions
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	 router.HandleFunc("/users-specific", controllers.SpecificDetails).Methods("GET")
	router.HandleFunc("/users", controllers.ReadUser).Methods("GET")
    router.HandleFunc("/users", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users",controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	 router.HandleFunc("/todo", controllers.CreateTodo).Methods("POST")

	
	// Start the HTTP server with the router
	log.Fatal(http.ListenAndServe(":8080", router))
}