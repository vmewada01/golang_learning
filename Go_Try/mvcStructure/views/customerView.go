package views

import (
	"encoding/json"
	"fmt"
	"mvcStructure/models"
	"net/http"
)

func RenderCustomerTemplate(w http.ResponseWriter, user *models.CustomerStructure) {
   // Encode user data as JSON
   jsonData, err := json.Marshal(user)
   if err != nil {
	   http.Error(w, "Error encoding user data to JSON", http.StatusInternalServerError)
	   return
   }

   fmt.Println("created customer response",string(jsonData))
   // Set content type and write JSON response
   w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(http.StatusOK)
   w.Write(jsonData)
}
