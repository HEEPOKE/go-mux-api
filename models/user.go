package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
	Tel      string `json:"tel"`
	Role     int    `json:"role"`
	Shop     []Shop
}
