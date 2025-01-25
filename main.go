package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iczky/delis-backend/shared"
)

func main() {
	config := shared.LoadConfig()

	shared.ConnectDatabase(config)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return shared.SuccessResponse(c, fiber.StatusOK, "Hello World", "Hai testing air live reloading")
	})

	log.Printf("starting server on port %s", config.AppPort)
	log.Fatal(app.Listen(":" + config.AppPort))
}