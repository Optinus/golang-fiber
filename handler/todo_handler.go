package handler

import (
	"strconv"

	"github.com/Optinus/golang-fiber/database"
	"github.com/Optinus/golang-fiber/model"
	"github.com/gofiber/fiber/v2"
)

// CreateTodo creates a new todo
func CreateTodo(c *fiber.Ctx) error {
	// Mendeklarasikan variabel untuk menyimpan data todo baru
	var newTodo model.Todo

	// Mem-parsing body request dan mengecek apakah terdapat kesalahan
	if err := c.BodyParser(&newTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Gagal memparsing body")
	}

	// Menetapkan ID baru untuk todo berdasarkan panjang database.Todos
	newTodo.ID = len(database.Todos) + 1

	// Menambahkan todo baru ke dalam database.Todos
	database.Todos = append(database.Todos, newTodo)

	// Mengembalikan respons dalam bentuk JSON
	return c.JSON(newTodo)
}

// GetTodoById retrieves a todo by its ID
func GetTodoById(c *fiber.Ctx) error {
	// Mendapatkan ID dari parameter URL
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)

	// Memeriksa apakah ID dapat diubah menjadi angka
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("ID harus berupa angka")
	}

	// Mencari todo berdasarkan ID dalam database.Todos
	for _, todo := range database.Todos {
		if todo.ID == id {
			// Mengembalikan todo dalam bentuk JSON jika ditemukan
			return c.JSON(todo)
		}
	}

	// Mengembalikan pesan kesalahan jika todo tidak ditemukan
	return c.Status(fiber.StatusNotFound).SendString("Todo tidak ditemukan")
}

// UpdateTodo updates a todo by its ID
func UpdateTodo(c *fiber.Ctx) error {
	// Mendapatkan ID dari parameter URL
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)

	// Memeriksa apakah ID dapat diubah menjadi angka
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("ID harus berupa angka")
	}

	// Mendeklarasikan variabel untuk menyimpan data todo yang diperbarui
	var updatedTodo model.Todo

	// Mem-parsing body request dan mengecek apakah terdapat kesalahan
	if err := c.BodyParser(&updatedTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Gagal mem-parsing body request")
	}

	// Menetapkan ID untuk todo yang diperbarui
	updatedTodo.ID = id

	// Mencari todo berdasarkan ID dalam database.Todos
	for i, todo := range database.Todos {
		if todo.ID == id {
			// Memperbarui todo dalam database dan mengembalikan todo yang diperbarui dalam bentuk JSON
			database.Todos[i] = updatedTodo
			return c.JSON(updatedTodo)
		}
	}

	// Mengembalikan pesan kesalahan jika todo tidak ditemukan
	return c.Status(fiber.StatusNotFound).SendString("Todo tidak ditemukan")
}

// DeleteTodo deletes a todo by its ID
func DeleteTodo(c *fiber.Ctx) error {
	// Mendapatkan ID dari parameter URL
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)

	// Memeriksa apakah ID dapat diubah menjadi angka
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("ID harus berupa angka")
	}

	// Mencari todo berdasarkan ID dalam database.Todos
	for i, todo := range database.Todos {
		if todo.ID == id {
			// Menghapus todo dari database.Todos dan mengembalikan pesan berhasil dihapus
			database.Todos = append(database.Todos[:i], database.Todos[i+1:]...)
			return c.SendString("Todo berhasil dihapus")
		}
	}

	// Mengembalikan pesan kesalahan jika todo tidak ditemukan
	return c.Status(fiber.StatusNotFound).SendString("Todo tidak ditemukan")
}

// GetAllTodo retrieves all todos
func GetAllTodo(c *fiber.Ctx) error {
	// Mengembalikan seluruh data todo dalam database.Todos dalam bentuk JSON
	return c.JSON(database.Todos)
}
