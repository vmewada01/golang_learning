package views

import (
	"encoding/json"
	"mvcStructure/models"
	"net/http"
)

func RenderUserTemplate(w http.ResponseWriter, user *models.UserStructure) {
   // Encode user data as JSON
   jsonData, err := json.Marshal(user)
   if err != nil {
	   http.Error(w, "Error encoding user data to JSON", http.StatusInternalServerError)
	   return
   }

   // Set content type and write JSON response
   w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(http.StatusOK)
   w.Write(jsonData)
}
