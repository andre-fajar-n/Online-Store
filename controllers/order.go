package controllers

import (
	"net/http"

	"github.com/andre-fajar-n/Online-Store/models"
	"github.com/andre-fajar-n/Online-Store/services"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	StatusCode   int
	orderService services.OrderService
}

func (c *Controller) AddToCart(ctx *gin.Context) {
	req := new(models.CartRequest)
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"en": "Invalid request body",
			"id": "Request body tidak valid",
		})
		return
	}

	if err := c.orderService.AddToCart(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"en": "Internal server error",
			"id": "Terjadi kesalahan di server",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"en": "Success",
		"id": "Sukses",
	})
}
