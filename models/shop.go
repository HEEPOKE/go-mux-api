package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name    string `json:"name"`
	OwnerID uint   `json:"ownerID"`
}
