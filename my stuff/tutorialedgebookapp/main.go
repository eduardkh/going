package main

import (
	"fmt"
	"log"
	"tutorialedgebookapp/book"
	"tutorialedgebookapp/database"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()
	database.InitDatabase()

	// Migrate the schema
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("✔️  Database Migrated")
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
