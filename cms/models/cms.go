package models

import "time"

type MenuItem struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"desctiption"`
	ImageUrl    string    `json:"image_url"`
	CategoryId  uint      `json:"categoryId"`
	IsAvailable bool      `json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ItemCategory struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}