package repositories

import (
	"github.com/andre-fajar-n/Online-Store/config"
	"github.com/andre-fajar-n/Online-Store/models"
	"gorm.io/gorm"
)

const (
	StatusOrderCart = "CART"
)

type OrderRepo struct {
	conn *gorm.DB
}

func (r *OrderRepo) Create(data *models.CartRequest) error {
	r.conn = config.ConnDB
	var orderID uint
	dataOrder := models.Order{
		UserID:   data.CustomerID,
		StatusID: StatusOrderCart,
	}
	if err := r.conn.Create(&dataOrder).Error; err != nil {
		return err
	}

	dataOrderDetail := &models.OrderDetail{
		ProductID: data.ProductID,
		OrderID:   orderID,
		Quantity:  data.Quantity,
	}
	if err := r.conn.Create(dataOrderDetail).Error; err != nil {
		return err
	}

	return nil
}
