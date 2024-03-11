package controllers

import (
	"encoding/json"
	"io/ioutil"
	"mvcStructure/models"
	"mvcStructure/views"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    // Read request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

    // Parse JSON data into user struct
    var user models.UserStructure
    if err := json.Unmarshal(body, &user); err != nil {
        http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
        return
    }

    // Save user to the database
    // db.Create(&user)

    // Render user view
	views.RenderUserTemplate(w, &user)

}