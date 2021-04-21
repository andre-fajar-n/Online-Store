package repositories

import (
	"github.com/andre-fajar-n/Online-Store/config"
	"github.com/andre-fajar-n/Online-Store/models"
)

const (
	StatusOrderCart = "CART"
)

func CreateCart(data *models.CartRequest) error {
	conn := config.ConnDB
	var orderID uint
	dataOrder := models.Order{
		UserID:   data.CustomerID,
		StatusID: StatusOrderCart,
	}
	if err := conn.Create(&dataOrder).Error; err != nil {
		return err
	}

	dataOrderDetail := &models.OrderDetail{
		ProductID: data.ProductID,
		OrderID:   orderID,
		Quantity:  data.Quantity,
	}

	if err := conn.Create(dataOrderDetail).Error; err != nil {
		return err
	}

	return nil
}
