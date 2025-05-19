package handlers

import (
	//"fmt"
	"time"
	"github.com/KunachRatchapat/BookShop_GORM/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
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
		//Respone Success!
		return c.JSON(fiber.Map{
		"message": "Register Success !",
		})
		
	}
	
}

func LoginUser(db *gorm.DB, user *model.User) fiber.Handler{
	return func(c *fiber.Ctx) error {

		//Get User From Email
		selectUser := new(model.User)
		result := db.Where("email = ?" ,user.Email).First(selectUser)
		if result.Error != nil{
			return c.JSON(fiber.Map{
				"Error":"No Email In Database !",
			})
		}
		
		//comapare Password
		err := bcrypt.CompareHashAndPassword(
			[]byte(selectUser.Password), //From, Client
			[]byte(user.Password), //Inside DB
	)
		if err != nil {
			return c.JSON(fiber.Map{
				"Error":"Login Unsuccess !",
			}) 
		}

	// Create JWT
	token := jwt.New(jwt.SigningMethodHS256)
  	claims := token.Claims.(jwt.MapClaims)
  	claims["user_id"] = selectUser.ID
  	claims["exp"] = time.Now().Add(time.Hour * 720).Unix()

	t , err := token.SignedString([]byte(os.Getenv("jwtSecretKey")))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("No Token !")
		
	}
	return c.JSON(fiber.Map{
		"message":"Login Success",
		"token":t,
	})
  }
}