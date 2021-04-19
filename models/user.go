package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username uint   `json:"Username"`
	Password uint   `json:"password"`
	Role     string `gorm:"type:enum('ADMIN','CUSTOMER')" json:"role"`
}
