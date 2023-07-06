package controller

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mawitra/test/database"
	"github.com/mawitra/test/model/entity"
	"github.com/mawitra/test/model/request"
	"github.com/mawitra/test/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}
	log.Println(loginRequest)
	validate := validator.New()
	errvalidate := validate.Struct(loginRequest)
	if errvalidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"messagae": "failed",
			"error":    errvalidate.Error(),
		})
	}
	var user entity.User
	err := database.DB.First(&user, "username = ?", loginRequest.Username).Error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Username",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Password",
		})
	}

	claims := jwt.MapClaims{}
	claims["fullname"] = user.Fullname
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	token, errtoken := utils.GenerateToken(&claims)
	if errtoken != nil {
		log.Println(errtoken)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Credential",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
