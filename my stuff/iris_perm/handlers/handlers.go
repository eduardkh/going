package handlers

import (
	"time"

	"github.com/kataras/iris/v12"
)

var startTime = time.Now()

func Index(ctx iris.Context) {

	ctx.View("index.html", iris.Map{
		"Title":           "Page Title",
		"FooterText":      "Footer contents",
		"Message":         "Main contents",
		"Name":            "iris",
		"serverStartTime": startTime,
	})
}

func GetBooks(ctx iris.Context) {

	ctx.View("books.html", iris.Map{
		"Title":   "Page Title",
		"Message": "books name",
	})
}

func GetBook(ctx iris.Context) {

	ctx.View("book.html", iris.Map{
		"Title":   "Page Title",
		"Message": "ספר",
	})
}
