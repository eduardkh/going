package handlers

import (
	"time"

	"github.com/kataras/iris/v12"
)

var startTime = time.Now()
var year = startTime.Year()

func Index(ctx iris.Context) {

	ctx.View("index.html", iris.Map{
		"Title":           "Page Title",
		"Message":         "Main contents",
		"Name":            "iris",
		"serverStartTime": startTime,
		"FooterYear":      year,
	})
}

func GetBooks(ctx iris.Context) {

	ctx.View("books.html", iris.Map{
		"Title":      "Page Title",
		"Message":    "books name",
		"FooterYear": year,
	})
}

func GetBook(ctx iris.Context) {

	ctx.View("book.html", iris.Map{
		"Title":      "Page Title",
		"Message":    "ספר",
		"FooterYear": year,
	})
}
