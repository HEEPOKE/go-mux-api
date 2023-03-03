package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    uint   `json:"price"`
}
