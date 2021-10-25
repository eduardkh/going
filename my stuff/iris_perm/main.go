package main

import (
	"time"

	_ "github.com/iris-contrib/pongo2-addons/v4"
	"github.com/kataras/iris/v12"
)

var startTime = time.Now()

func main() {
	app := iris.New()

	tmpl := iris.Django("./templates", ".html")
	tmpl.Reload(true)                             // reload templates on each request (development mode)
	tmpl.AddFunc("greet", func(s string) string { // {{greet(name)}}
		return "Greetings " + s + "!"
	})

	app.RegisterView(tmpl)

	app.Get("/", hi)

	app.Listen(":8080")
}

func hi(ctx iris.Context) {

	ctx.View("hi.html", iris.Map{
		"title":           "Hi Page",
		"name":            "iris",
		"serverStartTime": startTime,
	})
}
