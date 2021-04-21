package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string        `json:"name"`
	Quantity     uint          `json:"quantity"`
	Price        uint          `json:"price"`
	OrderDetails []OrderDetail `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
