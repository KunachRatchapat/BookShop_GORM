package handlers

import (
	//"fmt"

	

	"github.com/KunachRatchapat/BookShop_GORM/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB)  fiber.Handler{
	return func (c *fiber.Ctx)  error{
		//Get Data User
		user := new(model.User)
		if err := c.BodyParser(user) ; err !=nil{
			return err
		}
		//Encrypt  the Password
		hashPassword ,err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
		if err != nil{
			return err
		}
		//change password to crazy string
		user.Password = string(hashPassword)

		//Create User
		result := db.Create(user)
		if result.Error != nil{
			return c.JSON(fiber.Map{
				"Error":"Unsuccess  Register !",
			})
		}

		return c.JSON((user) fiber.Map{
		"Success": "Register Success !",
		})
		
	}

	
}