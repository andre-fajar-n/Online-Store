package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `gorm:"type:enum('ADMIN','CUSTOMER')" json:"role"`
}

type UserCreate struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"Username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
