package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/go-webstuffs/setting-up-backend/db"
	"github.com/sidmohanty11/go-webstuffs/setting-up-backend/routes"
)

func main() {
	// server connection
	app := fiber.New()
	routes.Setup(app)

	// db connection
	db.Connect()

	// server listening port
	app.Listen(":8000")
}
