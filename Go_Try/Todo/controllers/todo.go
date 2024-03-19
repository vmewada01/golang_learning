package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vmewada01/todo/config"
	"github.com/vmewada01/todo/models"
	"github.com/vmewada01/todo/utils"
)

var NewTodo models.Todo 


func GetBook(w http.ResponseWriter, r *http.Request){
	newTodo := models.GetAllTodo()
	res ,_ := json.Marshal(newTodo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err != nil {
        fmt.Println("Error while parsing")
        return
    }
    todo, err := models.GetTodoByID(ID)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode("Todo not found")
        return
    }
    res, _ := json.Marshal(todo)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}


func CreateBook(w http.ResponseWriter, r *http.Request){
	makeTodo := &models.Todo{}
     utils.ParseBody(r, makeTodo)
    b := makeTodo.CreateTodo()
    res, _ := json.Marshal(b)
    w.WriteHeader(http.StatusOK)
    w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err!= nil {
        fmt.Println("Error while parsing")
            return
    }
	bookDetails:= models.DeleteTodo(ID)
    res,_ := json.Marshal(bookDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}


func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateTodo = &models.Todo{}
	utils.ParseBody(r, updateTodo)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
    ID, err := strconv.ParseInt(bookId, 0, 0)
    if err!= nil {
        fmt.Println("Error while parsing")
            return
    }
	todo, err := models.GetTodoByID(ID)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode("Todo not found")
        return
    }
	
    fmt.Println("******************", todo)
	if updateTodo.Title == "" {
		todo.Title = updateTodo.Title
	}
	if updateTodo.Description == "" {
        todo.Description = updateTodo.Description
    }
	if updateTodo.Content == "" {
        todo.Content = updateTodo.Content
    }
	 fmt.Println("---------------",todo)
     config.GetDB().Save(&todo) 
	 res, _ := json.Marshal(todo)
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusOK)
     w.Write(res)
}