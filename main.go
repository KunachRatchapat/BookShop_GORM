package main

import (
	
	"log"
	"github.com/KunachRatchapat/BookShop_GORM/database"
	"github.com/KunachRatchapat/BookShop_GORM/model"
	"github.com/KunachRatchapat/BookShop_GORM/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	
)

func main(){
	//Load Envfile
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error Loading Env !")
	}

	db := database.ConnectDB()
	db.AutoMigrate(&model.Book{})

	//Setup Server
	app := fiber.New()

	//Connect Router
	routes.RouteInit(app,db)
	//Port
	app.Listen(":5000")



}