package main

import (
	"fmt"
	"log"
	"tutorialedgebookapp/book"
	"tutorialedgebookapp/database"
	"tutorialedgebookapp/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitDatabase()

	// Migrate the schema
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("✔️  Database Migrated")
	routes.BookApiRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
