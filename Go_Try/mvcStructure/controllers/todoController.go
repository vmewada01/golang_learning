package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mvcStructure/initializers"
	"mvcStructure/models"
	"mvcStructure/views"
	"net/http"

	"github.com/google/uuid"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Parse JSON data into user struct
	var todo models.TodosModel
	if err := json.Unmarshal(body, &todo); err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
		return
	}

	// Generate UUID
	todo.UUID = uuid.New().String()

	// Save todo to the database
	if err := initializers.DB.Create(&todo).Error; err != nil {
		http.Error(w, "Error saving todo to database", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Render todo view
	views.RenderTodosTemplate(w, &todo)

}
