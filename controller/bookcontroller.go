package controller

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mawitra/test/database"
	"github.com/mawitra/test/model/entity"
	"github.com/mawitra/test/model/request"
)

func BookIndex(c *fiber.Ctx) error {
	var books []entity.Book
	err := database.DB.Find(&books).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(fiber.Map{
		"data": books,
	})
}

func BookShow(c *fiber.Ctx) error {
	ID := c.Params("id")
	var books entity.Book
	err := database.DB.First(&books, "id = ?", ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Data Not Found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "berhasil",
		"data":    books,
	})
}

func BookCreate(c *fiber.Ctx) error {
	bookvalidate := new(request.BookRequest)
	if err := c.BodyParser(bookvalidate); err != nil {
		return err
	}

	validate := validator.New()
	errvalidate := validate.Struct(bookvalidate)
	if errvalidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"messagae": "failed",
			"error":    errvalidate,
		})
	}
	bookcreate := entity.Book{
		Title:     bookvalidate.Title,
		Published: bookvalidate.Published,
		Isbn:      bookvalidate.Isbn,
	}

	errCreateBook := database.DB.Create(&bookcreate).Error
	if errCreateBook != nil {
		return c.Status(500).JSON(fiber.Map{
			"messagae": "failed to store data",
		})
	}
	return c.JSON(fiber.Map{
		"messagae": "Created Data Successfully",
		"data":     bookcreate,
	})

}

func BookUpdate(c *fiber.Ctx) error {
	bookRequest := new(request.BookRequestUpdate)
	if err := c.BodyParser(bookRequest); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	var book entity.Book

	ID := c.Params("id")
	err := database.DB.First(&book, "id = ?", ID).Error
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

	book.Title = bookRequest.Title
	book.Published = bookRequest.Published
	book.Isbn = bookRequest.Isbn

	errupdate := database.DB.Save(&book).Error
	if errupdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"messagae": "Internal Server Error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Updated Data Successfully",
		"data":    book,
	})

}

func BookDelete(c *fiber.Ctx) error {
	ID := c.Params("id")
	var book entity.Book

	err := database.DB.Debug().First(&book, ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"messagae": "User Not Found",
		})
	}

	errDelete := database.DB.Debug().Delete(&book).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"messagae": "Internal Server Error",
		})
	}
	return c.JSON(fiber.Map{
		"messagae": "data Deleted",
	})

}
