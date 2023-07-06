package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mawitra/test/controller"
	"github.com/mawitra/test/middleware"
)

func Router(r *fiber.App) {
	r.Post("/login", controller.Login)
	r.Post("/register", controller.UserCreate)

	r.Get("/user", middleware.UserMiddleware, controller.UserIndex)
	r.Get("/user/:id", controller.UserShow)
	r.Put("/user/:id", controller.UserUpdate)
	r.Delete("/user/:id", controller.UserDelete)

	r.Get("/book", controller.BookIndex)
	r.Post("/book", controller.BookCreate)
	r.Get("/book/:id", controller.BookShow)
	r.Put("/book/:id", controller.BookUpdate)
	r.Delete("/book/:id", controller.BookDelete)

	r.Get("/author", controller.AuthorIndex)
	r.Post("/author", controller.AuthorCreate)
	r.Get("/author/:id", controller.AuthorShow)
	r.Put("/author/:id", controller.AuthorUpdate)
	r.Delete("/author/:id", controller.AuthorDelete)

}
