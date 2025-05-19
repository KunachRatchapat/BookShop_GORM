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

//Update Book
func UpdateBooks(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get ID from params
		id := c.Params("id")

		// Convert to int
		bookId, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID!",
			})
		}

		// Find book by ID
		var book model.Book
		if err := db.First(&book, bookId).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found!",
			})
		}

		// Parse update data
		newData := new(model.Book)
		if err := c.BodyParser(newData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid data!",
			})
		}

		// Update fields
		book.Author = newData.Author
		book.Description = newData.Description
		book.Price = newData.Price
		book.Publisher = newData.Publisher

		// Save changes
		if err := db.Save(&book).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save book!",
			})
		}

		// Return updated book
		return c.JSON(book)
	}
}


//Delete Book
func DeleteBook(db *gorm.DB) fiber.Handler{
	return func(c *fiber.Ctx) error{
		id,err := strconv.Atoi(c.Params("id"))
		if err != nil{
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"Error":"Invalid Id !",
			})
		}

		db.Delete(&model.Book{}, id)
		return c.SendString("Delete Success !!")
	}
}








