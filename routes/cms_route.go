package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iczky/delis-backend/cms/handler"
)

func RegisterCMSRoutes(app *fiber.App, cmsHandler *handler.CMSHandler) {
	api := app.Group("api/v1")

	api.Get("/menu-items", cmsHandler.GetAllMenuItems)
	api.Get("/menu-items/:id", cmsHandler.GetMenuItemByID)
	api.Post("/menu-items", cmsHandler.CreateMenuItem)
	api.Put("/menu-items", cmsHandler.UpdateMenuItem)
	api.Delete("/menu-items/:id", cmsHandler.DeleteMenuItem)

	api.Get("/categories", cmsHandler.GetAllCategories)
	api.Post("/categories", cmsHandler.CreateCategory)
}