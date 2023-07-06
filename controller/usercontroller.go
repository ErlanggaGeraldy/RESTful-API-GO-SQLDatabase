package controller

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mawitra/test/database"
	"github.com/mawitra/test/model/entity"
	"github.com/mawitra/test/model/request"
	"github.com/mawitra/test/utils"
)

func UserIndex(c *fiber.Ctx) error {
	var users []entity.User
	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(fiber.Map{
		"data": users,
	})
}

func UserShow(c *fiber.Ctx) error {
	ID := c.Params("id")
	var user entity.User
	err := database.DB.First(&user, "id = ?", ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil",
		"data":    user,
	})
}

func UserCreate(c *fiber.Ctx) error {
	var user entity.User
	uservalidate := new(request.UserCreateRequest)
	if err := c.BodyParser(uservalidate); err != nil {
		return err
	}

	errusername := database.DB.First(&user, "username = ?", uservalidate.Username).Error
	if errusername == nil {
		return c.Status(402).JSON(fiber.Map{
			"message": "Username already used",
		})
	}

	validate := validator.New()
	errvalidate := validate.Struct(uservalidate)
	if errvalidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"messagae": "failed",
			"error":    errvalidate,
		})
	}
	newUser := entity.User{
		Fullname: uservalidate.Fullname,
		Username: uservalidate.Username,
	}

	hashedPassword, err := utils.HashingPassword(uservalidate.Password)
	hashedPassword1, err := utils.HashingPassword(uservalidate.RepeatPassword)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"messagae": "Internal Server Eror",
		})
	}
	newUser.Password = hashedPassword
	newUser.RepeatPassword = hashedPassword1

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"messagae": "failed to store data",
		})
	}
	return c.JSON(fiber.Map{
		"messagae": "Created Data Successfully",
		"data":     newUser,
	})

}

func UserUpdate(c *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	var user entity.User

	ID := c.Params("id")
	err := database.DB.First(&user, "id = ?", ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}
	hashedPassword, err := utils.HashingPassword(userRequest.Password)
	hashedPassword1, err := utils.HashingPassword(userRequest.RepeatPassword)

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"messagae": "Internal Server Eror",
		})
	}

	user.Username = userRequest.Username
	user.Password = hashedPassword
	user.RepeatPassword = hashedPassword1
	errupdate := database.DB.Save(&user).Error
	if errupdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"messagae": "Internal Server Error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Updated Data Successfully",
		"data":    user,
	})

}

func UserDelete(c *fiber.Ctx) error {
	ID := c.Params("id")
	var user entity.User

	err := database.DB.Debug().First(&user, ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"messagae": "User Not Found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"messagae": "Internal Server Error",
		})
	}
	return c.JSON(fiber.Map{
		"messagae": "User Deleted",
	})

}
