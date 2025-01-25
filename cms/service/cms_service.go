package service

import (
	"github.com/iczky/delis-backend/cms/models"
	"github.com/iczky/delis-backend/cms/repository"
)

type CMSService interface {
	GetAllMenuItems() ([]models.MenuItem, error)
	GetMenuItemByID(id uint) (*models.MenuItem, error)
	CreateMenuItem(item *models.MenuItem) error
	UpdateMenuItem(item *models.MenuItem) error
	DeleteMenuItem(id uint) error

	GetAllCategories() ([]models.ItemCategory, error)
	GetCategoryByID(id uint) (*models.ItemCategory, error)
	CreateCategory(category *models.ItemCategory) error
	UpdateCategory(category *models.ItemCategory) error
	DeleteCategory(id uint) error
}

type cmsService struct {
	repo repository.CMSRepository
}

func NewCMSService(repo repository.CMSRepository) CMSService {
	return &cmsService{repo: repo}
}

// CreateCategory implements CMSService.
func (c *cmsService) CreateCategory(category *models.ItemCategory) error {
	return c.repo.CreateCategory(category)
}

// CreateMenuItem implements CMSService.
func (c *cmsService) CreateMenuItem(item *models.MenuItem) error {
	return c.repo.CreateMenuItem(item)
}

// DeleteCategory implements CMSService.
func (c *cmsService) DeleteCategory(id uint) error {
	return c.repo.DeleteCategory(id)
}

// DeleteMenuItem implements CMSService.
func (c *cmsService) DeleteMenuItem(id uint) error {
	return c.repo.DeleteMenuItem(id)
}

// GetAllCategories implements CMSService.
func (c *cmsService) GetAllCategories() ([]models.ItemCategory, error) {
	return c.repo.GetAllCategories()
}

// GetAllMenuItems implements CMSService.
func (c *cmsService) GetAllMenuItems() ([]models.MenuItem, error) {
	return c.repo.GetAllMenuItems()
}

// GetCategoryByID implements CMSService.
func (c *cmsService) GetCategoryByID(id uint) (*models.ItemCategory, error) {
	return c.repo.GetCategoryByID(id)
}

// GetMenuItemByID implements CMSService.
func (c *cmsService) GetMenuItemByID(id uint) (*models.MenuItem, error) {
	return c.repo.GetMenuItemByID(id)
}

// UpdateCategory implements CMSService.
func (c *cmsService) UpdateCategory(category *models.ItemCategory) error {
	return c.repo.UpdateCategory(category)
}

// UpdateMenuItem implements CMSService.
func (c *cmsService) UpdateMenuItem(item *models.MenuItem) error {
	return c.repo.UpdateMenuItem(item)
}


