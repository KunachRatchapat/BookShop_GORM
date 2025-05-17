package handlers

import (
	//"fmt"
	"strconv"
	"github.com/KunachRatchapat/BookShop_GORM/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

//Create Book
func CreateBook(db *gorm.DB ) fiber.Handler{
	return func (c *fiber.Ctx) error{
		book := new(model.Book)
		if err := c.BodyParser(book);  err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error Adding Book !")
	}
		//Add NewBook in DB !
		db.Create(book)
		return c.JSON(book)
	}

}	

//Get 1 Book By ID
func GetsOneBook(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error{
		// Get Id Client
		id,err := strconv.Atoi(c.Params("id"))
		if err != nil{
		return c.Status(fiber.StatusBadGateway).SendString("Not Found ID !")
	}
	var book model.Book
	db.First(&book,id)
	
	//JSON Send to Client
	return c.JSON(book)
	}	
}

//Gets All Books
func GetAllBooks(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error{
	var book []model.Book
	db.Find(&book)

	//Send To Client
	return c.JSON(book)
	}
}	

//









