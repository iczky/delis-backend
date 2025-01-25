package repository

import (
	"github.com/iczky/delis-backend/cms/models"
	"gorm.io/gorm"
)

type CMSRepository interface {
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

type cmsRepository struct {
	db *gorm.DB
}

func NewCMSRepository(db *gorm.DB) CMSRepository {
	return &cmsRepository{db: db}
}
// CreateCategory implements CMSRepository.
func (c *cmsRepository) CreateCategory(category *models.ItemCategory) error {
	return c.db.Create(category).Error
}

// CreateMenuItem implements CMSRepository.
func (c *cmsRepository) CreateMenuItem(item *models.MenuItem) error {
	return c.db.Create(item).Error
}

// DeleteCategory implements CMSRepository.
func (c *cmsRepository) DeleteCategory(id uint) error {
	return c.db.Delete(&models.MenuItem{}, id).Error
}

// DeleteMenuItem implements CMSRepository.
func (c *cmsRepository) DeleteMenuItem(id uint) error {
	return c.db.Delete((&models.MenuItem{}), id).Error
}

// GetAllCategories implements CMSRepository.
func (c *cmsRepository) GetAllCategories() ([]models.ItemCategory, error) {
	var categories []models.ItemCategory
	err := c.db.Find(&categories).Error
	return categories, err
}

// GetAllMenuItems implements CMSRepository.
func (c *cmsRepository) GetAllMenuItems() ([]models.MenuItem, error) {
	var menuItems []models.MenuItem
	err := c.db.Find(&menuItems).Error
	return menuItems, err
}

// GetCategoryByID implements CMSRepository.
func (c *cmsRepository) GetCategoryByID(id uint) (*models.ItemCategory, error) {
	var category models.ItemCategory
	err := c.db.First(&category, id).Error
	return &category, err
}

// GetMenuItemByID implements CMSRepository.
func (c *cmsRepository) GetMenuItemByID(id uint) (*models.MenuItem, error) {
	var item models.MenuItem
	err := c.db.First(&item, id).Error
	return &item, err
}

// UpdateCategory implements CMSRepository.
func (c *cmsRepository) UpdateCategory(category *models.ItemCategory) error {
	return c.db.Save(category).Error
}

// UpdateMenuItem implements CMSRepository.
func (c *cmsRepository) UpdateMenuItem(item *models.MenuItem) error {
	return c.db.Save(item).Error
}

