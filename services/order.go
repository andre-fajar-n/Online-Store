package services

import (
	"fmt"

	"github.com/andre-fajar-n/Online-Store/models"
	"github.com/andre-fajar-n/Online-Store/repositories"
	"github.com/andre-fajar-n/Online-Store/services/order"
	"github.com/gin-gonic/gin"
)

const (
	StatusOrderCart = "CART"
)

type OrderService struct {
	orderRepo    repositories.OrderRepo
	orderProduct repositories.ProductRepo
}

func (s *OrderService) AddToCart(data *models.CartRequest) error {
	if err := order.Validate(&gin.Context{}, data); err != nil {
		return err
	}

	// check if productID is exist
	product, err := s.orderProduct.GetByID(data.ProductID)
	if err != nil {
		return err
	}

	// check if quantity product less than quantity request
	if product.Quantity < data.Quantity {
		return fmt.Errorf("data not found")
	}

	if err := s.orderRepo.Create(data); err != nil {
		return err
	}

	return nil
}
