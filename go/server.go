package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gofiber/fiber"
)

type User struct {
	Id int
	Name string
}

func setupRoutes(app *fiber.App) {
	app.Get("/", test.Test)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(10000)

	defer database.DBConn.Close()
}