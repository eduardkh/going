package main

import (
	"log"
	"login/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.Routes(app)
	log.Fatal(app.Listen(":3000"))
}
