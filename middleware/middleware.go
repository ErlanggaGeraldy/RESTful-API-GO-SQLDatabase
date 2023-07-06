package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mawitra/test/utils"
)

func UserMiddleware(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"messagae": "unauthenticated",
		})
	}

	_, err := utils.VerifyToken(token)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"messagae": "unauthenticated",
		})
	}
	return c.Next()
}
