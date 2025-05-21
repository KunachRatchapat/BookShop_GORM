package routes

import (
	"github.com/KunachRatchapat/BookShop_GORM/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



func RouteInit(r *fiber.App , db *gorm.DB) {

	//Endpoint CRUD Users
	r.Post("/addbooks",handlers.CreateBook(db))
	r.Get("/books/:id",handlers.GetsOneBook(db))
	r.Get("/books",handlers.GetAllBooks(db))
	r.Put("/books/:id",handlers.UpdateBooks(db))
	r.Delete("/books/:id",handlers.DeleteBook(db))
	
	//Endpoint Register
	r.Post("/register",handlers.CreateUser(db))

	//Endpoint Login
	r.Post("/login", handlers.LoginUser(db))


}
	

