package views

import (
	"encoding/json"
	"fmt"
	"mvcStructure/models"
	"net/http"
)

func RenderTodosTemplate(w http.ResponseWriter, todo *models.TodosModel) {
   // Encode todo data as JSON
   jsonData, err := json.Marshal(todo)
   if err != nil {
	   http.Error(w, "Error encoding user data to JSON", http.StatusInternalServerError)
	   return
   }

   fmt.Println("created todo response",string(jsonData))
   // Set content type and write JSON response
   w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(http.StatusOK)
   w.Write(jsonData)
}
