package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ramses2099/webapi/controllers"
	"github.com/ramses2099/webapi/database"
	"github.com/ramses2099/webapi/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//
func setupRouters(app *fiber.App) {
	app.Get("/api/v1/book", controllers.GetBooks)
	app.Get("/api/v1/book/:id", controllers.GetBook)
	app.Post("/api/v1/book", controllers.NewBooks)
	app.Delete("/api/v1/book/:id", controllers.DeleteBook)
}

//
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connection successfully opened")

	// Auto migration
	database.DBConn.AutoMigrate(&models.Book{})
	fmt.Println("Database Migrated")

}

//
func main() {
	app := fiber.New()
	// init database
	initDatabase()

	// set up routers
	setupRouters(app)
	// set port running api
	app.Listen(":3001")
}
