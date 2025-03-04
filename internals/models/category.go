package models

import "gorm.io/gorm"


type Category struct {
    gorm.Model
    Name string `json:"name"`
    Slug string `json:"slug"`
    Products []Product
}