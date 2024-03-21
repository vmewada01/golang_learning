package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type books struct {
	ID  string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}


var book = []books{
	{ID : "1", Title: "In search of lost time", Author: "Marcel Proust", Quantity: 1 },
	{ID : "2", Title: "In search of peace  lost time", Author: "Nelson Mandela", Quantity: 3 },
	{ID : "3", Title: "Devil in the ocean ", Author: "Mark Thunder", Quantity: 5 },
}


func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, book)
}


func createBook(c *gin.Context){
	var newBook books

   if err := c.BindJSON(&newBook); err != nil {
	    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
   }

     book =	append(book, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)


}

func main() {
	fmt.Println("Gin Routing testing part")

	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")



}