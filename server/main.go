package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

func main() {
	fmt.Println("Hello world")

	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{}))

	todos := []Todo{}
	app.Get("/Todo", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		todo.ID = len(todos) + 1

		todos = append(todos, *todo)
		return c.JSON(todos)
	})
	app.Patch("/api/todos/:id/status", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}
		for i, t := range todos {
			if t.ID == id {
				todos[i].Status = true
				break
			}
		}
		return c.JSON(todos)
	})
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}
		for i, t := range todos {
			if t.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				break
			}
		}
		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})
	log.Fatal(app.Listen(":4000"))
}
