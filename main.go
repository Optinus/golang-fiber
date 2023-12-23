package main

import (
	"github.com/Optinus/golang-fiber/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRouter(app)
	app.Listen(":3000")
}
