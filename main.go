package main

import (
	"backend-koperasi/database"
	"backend-koperasi/migration"
	"backend-koperasi/route"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initial Database
	database.DatabaseInit()

	// run migration setelah run database diatas
	migration.Migrate()

	// Initial Fiber
	app := fiber.New()

	// Initial Route
	route.RouteInit(app)

	// Listen App
	log.Fatal(app.Listen(":8080"))
}
