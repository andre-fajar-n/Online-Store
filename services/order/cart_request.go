package order

import (
	"github.com/andre-fajar-n/Online-Store/helpers"
	"github.com/andre-fajar-n/Online-Store/models"
)

func Validate(req *models.CartRequest) error {
	if req.CustomerID <= 0 {
		return helpers.ErrorValidation(&helpers.ErrorResponse{
			En: "CustomerID must be more than 0",
			Id: "CustomerID harus lebih dari 0",
		})
	}

	if req.ProductID <= 0 {
		return helpers.ErrorValidation(&helpers.ErrorResponse{
			En: "ProductID must be more than 0",
			Id: "ProductID harus lebih dari 0",
		})
	}

	if req.Quantity <= 0 {
		return helpers.ErrorValidation(&helpers.ErrorResponse{
			En: "Quantity must be more than 0",
			Id: "Quantity harus lebih dari 0",
		})
	}

	return nil
}
