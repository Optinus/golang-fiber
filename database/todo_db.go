package database

import "github.com/Optinus/golang-fiber/model"

// Todos adalah slice untuk menyimpan daftar tugas (todos)
var Todos = []model.Todo{
	{ID: 1, Title: "Belajar Golang", Description: "Fokus di dasar-dasar Golang saja", Finished: false},
	// Anda dapat menambahkan data todos lainnya di sini
}
