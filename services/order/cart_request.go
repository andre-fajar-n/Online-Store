package order

import (
	"net/http"

	"github.com/andre-fajar-n/Online-Store/models"
	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context, req *models.CartRequest) error {
	if req.CustomerID <= 0 {
		return http.ErrBodyNotAllowed
	}

	if req.ProductID <= 0 {
		return http.ErrBodyNotAllowed
	}

	if req.Quantity <= 0 {
		return http.ErrBodyNotAllowed
	}

	return nil
}
