package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Author      string
	Description string
	Price       uint
	Publisher   string
}

// Create Book
func createBook(db *gorm.DB, book *Book) {
	result := db.Create(book)

	if result.Error != nil {
		log.Fatalf("Error creating Book: %v", result.Error)
	}

	fmt.Println("Create Book Success !")
}

// Get Books
func getsBooks(db *gorm.DB, id uint) *Book {
	var book Book
	result := db.First(&book, id) //(variableinstruct,condition)

	if result.Error != nil {
		log.Fatalf("Error Showing Books:%v", result.Error)
	}

	return &book

}

// Update Books
func updateBooks(db *gorm.DB, book *Book) {
	result := db.Save(&book)

	if result.Error != nil {
		log.Fatalf("Update Don't Success !")
	}

	fmt.Println("Update Success")
}

// Delete Books
func deleteBooks(db *gorm.DB, id uint) {
	var book Book
	result := db.Delete(&book, id)
	if result.Error != nil {
		log.Fatal("Error deleting book: %v ", result.Error)
	}

	fmt.Println("Book Delete Successful !")
}

//SearchBooks
func searchBooks(db *gorm.DB, infobook string) *Book {
	var book Book

	result := db.Where("author = ?",infobook).First(&book)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound){	
			//Not fount Book
			return nil
		}
		//Error Connection 
		log.Fatalf("เกิดข้อผิดพลาด:%v",result.Error)
	}
	return &book

}
