package handler

import (
	"github.com/Optinus/golang-fiber/database"
	"github.com/Optinus/golang-fiber/model"
	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	var newTodo model.Todo

	if err := c.BodyParser(&newTodo); err != nil {
		return c.Status(404).SendString("Gagal memparsing body")
	}
	newTodo.ID = len(database.Todos) + 1

	database.Todos = append(database.Todos, newTodo)
	return c.JSON(newTodo)
}

func GetTodo(c *fiber.Ctx) error {
	return c.JSON(database.Todos)
}
