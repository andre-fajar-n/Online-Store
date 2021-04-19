package models

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	ProductID uint
	OrderID   uint
	Quantity  uint
	Product   []Product `gorm:"foreignKey:ID;references:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Order     []Order   `gorm:"foreignKey:ID;references:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
