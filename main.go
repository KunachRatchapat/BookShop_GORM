package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func main(){
	//Load Envfile
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error Loading Env !")
	}

	//Get From my env 
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	//Dsn For Connect
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host,port,user,password,dbname)

	// New logger for detailed SQL logging
  newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
    logger.Config{
      SlowThreshold: time.Second, // Slow SQL threshold
      LogLevel:      logger.Info, // Log level
      Colorful:      true,        // Enable color
    },
  )

	//Connect Database With GORM
	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err !=nil{
		panic("Failed to Connect Database !!")
	}
	print(db)
	
	db.AutoMigrate(&Book{})
	fmt.Println("Connect Success")

	//Start CreateBook
	/*newBooks := &Book{
		Author: "Harper Lee2",
		Description:"Racism and justice in the South" ,
		Price:  359,
		Publisher: "J.B. Lippincott & Co.",
	}

	//Use Func createBook
	createBook(db,newBooks) //Create INSERT(book) and Struct Book(newbooks)*/

	//Usefunc GetBoks
	/*currentbook := getsBooks(db,2)
	

	currentbook.Author = "J.K. Rowling Wizard2"
	updateBooks(db,currentbook)*/
	
	//Delete Func
	//deleteBooks(db,4)
	
	

}