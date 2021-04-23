package repositories

import (
	"fmt"

	"github.com/andre-fajar-n/Online-Store/config"
	"github.com/andre-fajar-n/Online-Store/models"
)

const (
	StatusOrderCart     = "CART"
	StatusOrderCheckout = "CHECKOUT"
)

func CreateCart(data *models.CartRequest) error {
	conn := config.ConnDB
	dataOrder := models.Order{
		UserID:   data.CustomerID,
		StatusID: StatusOrderCart,
	}
	if err := conn.Create(&dataOrder).Scan(&dataOrder).Error; err != nil {
		return err
	}

	dataOrderDetail := &models.OrderDetail{
		ProductID: data.ProductID,
		OrderID:   dataOrder.ID,
		Quantity:  data.Quantity,
	}
	fmt.Println("ORDER", dataOrder.ID)
	if err := conn.Create(dataOrderDetail).Error; err != nil {
		return err
	}

	return nil
}

func Checkout(data *models.CartRequest) error {
	conn := config.ConnDB

	dataOrder := models.Order{
		UserID:   data.CustomerID,
		StatusID: StatusOrderCheckout,
	}
	if err := conn.Create(&dataOrder).Scan(&dataOrder).Error; err != nil {
		return err
	}

	dataOrderDetail := &models.OrderDetail{
		ProductID: data.ProductID,
		OrderID:   dataOrder.ID,
		Quantity:  data.Quantity,
	}

	if err := conn.Create(dataOrderDetail).Error; err != nil {
		return err
	}

	return nil
}
