package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iczky/delis-backend/cms/handler"
	"github.com/iczky/delis-backend/cms/models"
	"github.com/iczky/delis-backend/cms/repository"
	"github.com/iczky/delis-backend/cms/service"
	"github.com/iczky/delis-backend/routes"
	"github.com/iczky/delis-backend/shared"
)

func main() {
	config := shared.LoadConfig()

	db := shared.ConnectDatabase(config)

	if err := db.AutoMigrate(&models.MenuItem{}, &models.ItemCategory{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	cmsRepo := repository.NewCMSRepository(db)
	cmsService := service.NewCMSService(cmsRepo)
	cmsHandler := handler.NewCMSHandler(cmsService)

	app := fiber.New()

	routes.RegisterCMSRoutes(app, cmsHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return shared.SuccessResponse(c, fiber.StatusOK, "Hello World", "Hai testing air live reloading")
	})

	log.Printf("starting server on port %s", config.AppPort)
	log.Fatal(app.Listen(":" + config.AppPort))
}