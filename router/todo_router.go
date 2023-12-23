package router

import (
	"github.com/Optinus/golang-fiber/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Post("/", handler.CreateTodo)
	app.Get("/", handler.GetTodo)
}
