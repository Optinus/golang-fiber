package main

import (
	"github.com/Optinus/golang-fiber/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Membuat objek aplikasi Fiber baru
	app := fiber.New()

	// Mengatur router dengan endpoint-endpoint yang ditentukan di package router
	router.SetupRouter(app)

	// Mendengarkan pada port 3000 untuk incoming HTTP requests
	app.Listen(":3000")
}
