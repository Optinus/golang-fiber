package router

import (
	"github.com/Optinus/golang-fiber/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRouter mengatur endpoint-endpoint untuk menangani HTTP requests
func SetupRouter(app *fiber.App) {
	// Menangani HTTP POST request untuk membuat tugas baru
	app.Post("/", handler.CreateTodo)

	// Menangani HTTP GET request untuk mendapatkan tugas berdasarkan ID
	app.Get("/:id", handler.GetTodoById)

	// Menangani HTTP PUT request untuk memperbarui tugas berdasarkan ID
	app.Put("/:id", handler.UpdateTodo)

	// Menangani HTTP DELETE request untuk menghapus tugas berdasarkan ID
	app.Delete("/:id", handler.DeleteTodo)

	// Menangani HTTP GET request untuk mendapatkan semua tugas
	app.Get("/", handler.GetAllTodo)
}
