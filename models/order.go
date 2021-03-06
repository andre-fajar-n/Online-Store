package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID   uint
	StatusID string `gorm:"type:enum('CART','CHECKOUT','PAID','EXPIRED','CANCELED')"`
	// User     []User `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CartRequest struct {
	ProductID  uint `json:"product_id"`
	CustomerID uint `json:"customer_id"`
	Quantity   uint `json:"quantity"`
}
