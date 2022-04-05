package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ramses2099/webapi/database"
	"github.com/ramses2099/webapi/models"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []models.Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book models.Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBooks(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book models.Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("No book found with given ID")
	}
	db.Delete(&book)
	return c.SendString("Book successfully deleted")
}
