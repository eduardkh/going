package book

import (
	"tutorialedgebookapp/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating string `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		return c.Status(404).SendString("No Book Found with ID: " + id)
	}
	return c.JSON(book)
}
func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(400).SendString("Bad Request")
	}
	db.Create(&book)
	return c.JSON(book)
}
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(404).SendString("No Book Found with ID: " + id)
	}
	db.Delete(&book, id)
	return c.JSON(book)
}
