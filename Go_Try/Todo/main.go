package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/vmewada01/todo/models"
)

func main() {
	fmt.Println("TODO APPLICATION")

    port := os.Getenv("MY_APP_PORT")
    if port == "" {
        port = "8080"
    }

    e := echo.New()
    // e.Use(middleware.Logger())
    // e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, "Well, hello there!")
    })

    // GET all Todo items
    e.GET("/todo", func(c echo.Context) error {
        todos := models.GetAllTodo()
        return c.JSON(http.StatusOK, todos)
    })

    // GET a Todo item by ID
    e.GET("/todo/:id", func(c echo.Context) error {
		ID := c.Param("id")
		id, err := strconv.ParseInt(ID, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID format")
		}
		todo, err := models.GetTodoByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, "Todo not found")
		}
        return c.JSON(http.StatusOK, todo)
    })

    // POST a new Todo item
    e.POST("/todo/create", func(c echo.Context) error {
        todo := new(models.Todo)
        if err := c.Bind(todo); err != nil {
            return err
        }
        createdTodo := todo.CreateTodo()
        return c.JSON(http.StatusCreated, createdTodo)
    })

    // DELETE a Todo item by ID
	e.DELETE("/todo/:id", func(c echo.Context) error {
		ID := c.Param("id")
		id, err := strconv.ParseInt(ID, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID format")
		}
	
		if err := models.DeleteTodo(id); err != nil {
			return c.JSON(http.StatusNotFound, "Todo not found")
		}
	
		return c.JSON(http.StatusOK, "Todo deleted successfully")
	})


  

    e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
    e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
    
}
