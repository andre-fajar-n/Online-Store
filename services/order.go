package services

import (
	"github.com/andre-fajar-n/Online-Store/helpers"
	"github.com/andre-fajar-n/Online-Store/models"
	"github.com/andre-fajar-n/Online-Store/repositories"
	"github.com/andre-fajar-n/Online-Store/services/order"
)

const (
	StatusOrderCart = "CART"
)

type OrderService struct {
	orderRepo    repositories.OrderRepo
	orderProduct repositories.ProductRepo
}

func (s *OrderService) AddToCart(data *models.CartRequest) error {
	if err := order.Validate(data); err != nil {
		return err
	}

	// check if productID is exist
	product, err := s.orderProduct.GetByID(data.ProductID)
	if err != nil {
		return helpers.ErrorValidation(&helpers.ErrorResponse{
			En: "Invalid productID",
			Id: "ProductID tidak valid",
		})
	}

	// check if quantity product less than quantity request
	if product.Quantity < data.Quantity {
		if product.Quantity == 0 {
			return helpers.ErrorValidation(&helpers.ErrorResponse{
				En: "This product has run out of stock",
				Id: "Produk ini sudah habis",
			})
		}

		return helpers.ErrorValidation(&helpers.ErrorResponse{
			En: "Your demand exceeds stock of this product",
			Id: "Permintaan Anda melebihi stok produk ini",
		})
	}

	if err := s.orderRepo.Create(data); err != nil {
		return err
	}

	return nil
}
