package routes

import (
	"login/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", handlers.Index)
}
