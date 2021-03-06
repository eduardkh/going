package routes

import (
	"tutorialedgebookapp/book"

	"github.com/gofiber/fiber/v2"
)

func BookApiRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
	app.Put("/api/v1/book/:id", book.UpdateBook)
}
