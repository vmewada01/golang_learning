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

    
    // Generate UUID
    user.UUID = uuid.New().String()

    // Save user to the database
    if err := initializers.DB.Create(&user).Error; err != nil {
        http.Error(w, "Error saving user to database", http.StatusInternalServerError)
        fmt.Println(err)
        return
    }


    // Render user view
	views.RenderUserTemplate(w, &user)

}


func ReadUser(w http.ResponseWriter, r *http.Request) {
    // Query the database to fetch all users
    var users []models.UserStructure
    if err := initializers.DB.Find(&users).Error; err != nil {
        http.Error(w, "Error fetching users from database", http.StatusInternalServerError)
        return
    }

    // Encode user data into JSON
    jsonData, err := json.Marshal(users)
    if err != nil {
        http.Error(w, "Error encoding user data to JSON", http.StatusInternalServerError)
        return
    }

    // Set content type and write JSON response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}


func UpdateUser(w http.ResponseWriter, r *http.Request) {
    // Read request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

    // Decode JSON data into user struct
    var updatedUser models.UserStructure
    if err := json.Unmarshal(body, &updatedUser); err != nil {
        http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
        return
    }

    // Extract user ID from query parameters
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Missing user ID", http.StatusBadRequest)
        return
    }

    // Validate and parse UUID
    userID, err := uuid.Parse(id)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Update user data where UUID matches
    result := initializers.DB.Model(&models.UserStructure{}).Where("uuid = ?", userID).Updates(updatedUser)

    // Check for errors
    if result.Error != nil {
        http.Error(w, "Error updating user", http.StatusInternalServerError)
        return
    }

    // Check if any rows were affected
    if result.RowsAffected == 0 {
        http.Error(w, "No users were updated", http.StatusNotFound)
        return
    }

    // Send success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User updated successfully"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from query parameters
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Missing user ID", http.StatusBadRequest)
        return
    }

    // Validate and parse UUID
    userID, err := uuid.Parse(id)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Delete user from the database where UUID matches
    result := initializers.DB.Where("uuid = ?", userID).Delete(&models.UserStructure{})

    // Check for errors
    if result.Error != nil {
        http.Error(w, "Error deleting user", http.StatusInternalServerError)
        return
    }

    // Check if any rows were affected
    if result.RowsAffected == 0 {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Send success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User deleted successfully"))
}


func SpecificDetails(w http.ResponseWriter, r *http.Request){
  // Query the database to fetch selected fields for all users
  var users []map[string]interface{}
  if err := initializers.DB.Model(&models.UserStructure{}).Select("name", "email", "uuid", "address", "phone").Find(&users).Error; err != nil {
      http.Error(w, "Error fetching users from database", http.StatusInternalServerError)
      return
  }

  // Encode user data into JSON
  jsonData, err := json.Marshal(users)
  if err != nil {
      http.Error(w, "Error encoding user data to JSON", http.StatusInternalServerError)
      return
  }

  // Set content type and write JSON response
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonData)
}