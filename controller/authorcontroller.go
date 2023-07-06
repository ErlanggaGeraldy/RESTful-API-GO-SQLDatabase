package controller

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mawitra/test/database"
	"github.com/mawitra/test/model/entity"
	"github.com/mawitra/test/model/request"
)

func AuthorIndex(c *fiber.Ctx) error {
	var author []entity.Author
	err := database.DB.Find(&author).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(fiber.Map{
		"data": author,
	})
}

func AuthorShow(c *fiber.Ctx) error {
	ID := c.Params("id")
	var author entity.Author
	err := database.DB.First(&author, "id = ?", ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil",
		"data":    author,
	})
}

func AuthorCreate(c *fiber.Ctx) error {
	authorvalidate := new(request.AuthorRequest)
	if err := c.BodyParser(authorvalidate); err != nil {
		return err
	}

	validate := validator.New()
	errvalidate := validate.Struct(authorvalidate)
	if errvalidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"messagae": "failed",
			"error":    errvalidate,
		})
	}
	authorcreate := entity.Author{
		Name:    authorvalidate.Name,
		Country: authorvalidate.Country,
	}

	errCreateAuthor := database.DB.Create(&authorcreate).Error
	if errCreateAuthor != nil {
		return c.Status(500).JSON(fiber.Map{
			"messagae": "failed to store data",
		})
	}
	return c.JSON(fiber.Map{
		"messagae": "Created Data Successfully",
		"data":     authorcreate,
	})

}

func AuthorUpdate(c *fiber.Ctx) error {
	authorRequest := new(request.AuthorRequestUpdate)
	if err := c.BodyParser(authorRequest); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	var author entity.Author

	ID := c.Params("id")
	err := database.DB.First(&author, "id = ?", ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"messagae": "Internal Server Eror",
		})
	}

	author.Name = authorRequest.Name
	author.Country = authorRequest.Country

	errupdate := database.DB.Save(&author).Error
	if errupdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"messagae": "Internal Server Error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Updated Data Successfully",
		"data":    author,
	})

}

func AuthorDelete(c *fiber.Ctx) error {
	ID := c.Params("id")
	var author entity.Author

	err := database.DB.Debug().First(&author, ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"messagae": "User Not Found",
		})
	}

	errDelete := database.DB.Debug().Delete(&author).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"messagae": "Internal Server Error",
		})
	}
	return c.JSON(fiber.Map{
		"messagae": "data Deleted",
	})

}
