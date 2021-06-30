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

func AddToCart(data *models.CartRequest) error {
	if err := order.Validate(data); err != nil {
		return err
	}

	// check if productID is exist
	product, err := repositories.GetOneProductByID(data.ProductID)
	if err != nil {
		return helpers.ErrorValidation(&helpers.Response{
			En: "Invalid productID",
			Id: "ProductID tidak valid",
		})
	}

	// check if quantity product less than quantity request
	if product.Quantity < data.Quantity {
		if product.Quantity == 0 {
			return helpers.ErrorValidation(&helpers.Response{
				En: "This product has run out of stock",
				Id: "Produk ini sudah habis",
			})
		}

		return helpers.ErrorValidation(&helpers.Response{
			En: "Your demand exceeds stock of this product",
			Id: "Permintaan Anda melebihi stok produk ini",
		})
	}

	if err := repositories.CreateCart(data); err != nil {
		return err
	}

	return nil
}

func Checkout(data *models.CartRequest) error {
	if err := order.Validate(data); err != nil {
		return err
	}

	// check if productID is exist
	_, err := repositories.GetOneProductByID(data.ProductID)
	if err != nil {
		return helpers.ErrorValidation(&helpers.Response{
			En: "Invalid productID",
			Id: "ProductID tidak valid",
		})
	}

	availProduct, err := repositories.CountProductAvailable(data.ProductID, data.CustomerID)
	if err != nil {
		return helpers.ErrorBadRequest(&helpers.Response{
			En: err.Error(),
			Id: err.Error(),
		})
	}

	if availProduct < int64(data.Quantity) {
		return helpers.ErrorValidation(&helpers.Response{
			En: "The stock of this product is lacking or empty",
			Id: "Persediaan produk ini tidak cukup atau habis",
		})
	}

	if err := repositories.Checkout(data); err != nil {
		return err
	}

	return nil
}
