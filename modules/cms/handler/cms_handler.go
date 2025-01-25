package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iczky/delis-backend/modules/cms/models"
	"github.com/iczky/delis-backend/modules/cms/service"
	"github.com/iczky/delis-backend/shared"
)

type CMSHandler struct {
	service service.CMSService
}

func NewCMSHandler(service service.CMSService) *CMSHandler {
	return &CMSHandler{service: service}
}

func (h *CMSHandler) GetAllMenuItems(c *fiber.Ctx) error {
	items, err := h.service.GetAllMenuItems()
	if err != nil {
		return shared.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return shared.SuccessResponse(c, fiber.StatusOK, "Success", items)
}

func (h *CMSHandler) GetMenuItemByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return shared.ErrorResponse(c, fiber.StatusBadRequest, "Invalid ID")
	}
	item, err := h.service.GetMenuItemByID(uint(id))
	if err != nil {
		return shared.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return shared.SuccessResponse(c, fiber.StatusOK, "Success", item)
}

func (h *CMSHandler) CreateMenuItem(c *fiber.Ctx) error {
	var item models.MenuItem
	if err := c.BodyParser(&item); err != nil {
		return shared.ErrorResponse(c, fiber.StatusBadRequest, "Invalid Input")
	}
	if err := h.service.CreateMenuItem(&item); err != nil {
		return shared.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return shared.SuccessResponse(c, fiber.StatusOK, "Item Created", item)
}

func (h *CMSHandler) UpdateMenuItem(c *fiber.Ctx) error {
	var item models.MenuItem
	if err := c.BodyParser(&item); err != nil {
		return shared.ErrorResponse(c, fiber.StatusBadRequest, "Invalid Input")
	}
	if err := h.service.UpdateMenuItem(&item); err != nil {
		return shared.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return shared.SuccessResponse(c, fiber.StatusOK, "Item Updated", item)
}

func (h *CMSHandler) DeleteMenuItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return shared.ErrorResponse(c, fiber.StatusBadRequest, "Invalid Id")
	}
	if err := h.service.DeleteMenuItem(uint(id)); err != nil {
		return shared.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return shared.SuccessResponse(c, fiber.StatusOK, "Item Deleted", nil)
}

// Category Handlers
func (h *CMSHandler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		return shared.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return shared.SuccessResponse(c, fiber.StatusOK, "Success", categories)
}

func (h *CMSHandler) CreateCategory(c *fiber.Ctx) error {
	var category models.ItemCategory
	if err := c.BodyParser(&category); err != nil {
		return shared.ErrorResponse(c, fiber.StatusBadRequest, "Invalid Input")
	}
	if err := h.service.CreateCategory(&category); err != nil {
		return shared.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return shared.SuccessResponse(c, fiber.StatusOK, "Category Created", category)
}