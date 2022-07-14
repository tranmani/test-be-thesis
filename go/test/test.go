package test

import (
	"fmt"

	"github.com/elliotforbes/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	ID string `json:"id"`
	Name  string `json:"name"`
}

func Test(c *fiber.Ctx) {
	db := database.DBConn
	var data []User
	
	db.Find("Select * from test").Scan(&data)
	
	c.JSON(data)
}