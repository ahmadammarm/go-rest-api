package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name"`
    Description string `json:"description"`
	Price int    `json:"price"`
    Slug string `json:"slug"`
    CategoryID int `json:"category_id"`
    Category Category

}
