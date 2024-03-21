package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
		
		app := fiber.New()

	
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})
	
		app.Get("/users/:id", func(c *fiber.Ctx) error {
			id := c.Params("id")
			return c.SendString("User ID: " + id)
		})
	
		app.Post("/users", func(c *fiber.Ctx) error {
			
			type User struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}
			var user User
			if err := c.BodyParser(&user); err != nil {
				return err
			}
			
			return c.JSON(user)
		})
	
		
		app.Listen(":7000")
}
