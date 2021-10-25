package main

import (
	"iris_perm/handlers"

	_ "github.com/iris-contrib/pongo2-addons/v4"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	templates := iris.Django("./templates", ".html")
	app.RegisterView(templates)

	// routes
	app.Get("/", handlers.Index)
	app.Get("/books", handlers.GetBooks)
	app.Get("/book", handlers.GetBook)

	app.Listen(":8080")
}
