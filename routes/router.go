package routes

import (
	"github.com/KunachRatchapat/BookShop_GORM/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



func RouteInit(r *fiber.App , db *gorm.DB) {
	
	r.Post("/addbooks",handlers.CreateBook(db))
	r.Get("/books/:id",handlers.GetsOneBook(db))
	r.Get("/books",handlers.GetAllBooks(db))


}
	

