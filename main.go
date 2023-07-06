package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mawitra/test/database"
	"github.com/mawitra/test/database/migration"
	"github.com/mawitra/test/route"
)

func main() {
	database.ConnectDatabase()
	migration.RunMigration()
	app := fiber.New()
	route.Router(app)

	app.Listen(":3000")
}
