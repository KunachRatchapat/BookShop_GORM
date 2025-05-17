package model

import (
	
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Publisher   string `json:"publisher"`
}






// Create Book
/*func createBook(db *gorm.DB, book *Book) {
	result := db.Create(book)

	if result.Error != nil {
		log.Fatalf("Error creating Book: %v", result.Error)
	}

	fmt.Println("Create Book Success !")
}

// Get Book 1
func getsBookOne(db *gorm.DB, id int) *Book {
	var book Book
	result := db.First(&book, id) //(variableinstruct,condition)

	if result.Error != nil {
		log.Fatalf("Error Showing Books:%v", result.Error)
	}

	return &book

}

//Get Book All
func getsBooks(db *gorm.DB) []Book {
	var books []Book
	result := db.Find(&books) //(variableinstruct,condition)

	if result.Error != nil {
		log.Fatalf("Error Showing Books:%v", result.Error)
	}

	return books

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
	//Found BookID
	if result.RowsAffected == 0 {
		fmt.Println("Not Found ID BOOK")
	}

	fmt.Println("Book Delete Successful !")
}

//SearchBooksฺBy AuthorName
func searchBookByAuthor(db *gorm.DB, authorbook string) []Book {
	var book []Book

	result := db.Where("author = ?",authorbook).Order("price").Find(&book)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound){	
			//Not fount Book
			return nil
		}
		//Error Connection 
		log.Fatalf("เกิดข้อผิดพลาด:%v",result.Error)
	}
	return book 

}

//SearchBooksBy Price
func searchBooksByPrice(db *gorm.DB, pricebook string) []Book {
	var book []Book

	result := db.Order("price").Find(&book)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound){	
			//Not fount Book
			return nil
		}
		//Error Connection 
		log.Fatalf("เกิดข้อผิดพลาด:%v",result.Error)
	}
	return book 

}*/


