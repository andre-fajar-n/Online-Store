package controllers

import (
	"github.com/andre-fajar-n/Online-Store/helpers"
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
		helpers.DefaultErrorBadRequest(ctx, helpers.InvalidRequestBody)
		return
	}

	if err := c.orderService.AddToCart(req); err != nil {
		helpers.ReturnError(ctx, err)
		return
	}

	helpers.SuccessCreate(ctx, nil)
}
