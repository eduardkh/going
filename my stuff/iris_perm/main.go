package main

import (
	"time"

	_ "github.com/iris-contrib/pongo2-addons/v4"
	"github.com/kataras/iris/v12"
)

var startTime = time.Now()

func main() {
	app := iris.New()

	templates := iris.Django("./templates", ".html")
	app.RegisterView(templates)

	app.Get("/", index)

	app.Listen(":8080")
}

func index(ctx iris.Context) {

	ctx.View("index.html", iris.Map{
		"Title":           "Page Title",
		"FooterText":      "Footer contents",
		"Message":         "Main contents",
		"Name":            "iris",
		"serverStartTime": startTime,
	})
}
