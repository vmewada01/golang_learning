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

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Parse JSON data into user struct
	var customer models.CustomerStructure
	if err := json.Unmarshal(body, &customer); err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
		return
	}

	// Generate UUID
	customer.UUID = uuid.New().String()

	// Save customer to the database
	if err := initializers.DB.Create(&customer).Error; err != nil {
		http.Error(w, "Error saving customer to database", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Render customer view
	views.RenderCustomerTemplate(w, &customer)

}
