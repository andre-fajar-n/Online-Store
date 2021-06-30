package controllers

import (
	"github.com/andre-fajar-n/Online-Store/helpers"
	"github.com/andre-fajar-n/Online-Store/models"
	"github.com/andre-fajar-n/Online-Store/services"
	"github.com/gin-gonic/gin"
)

// AddToCart godoc
// @Summary Add product to cart
// @Description Add product to cart
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body models.CartRequest true "Cart request"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Router /cart [post]
func AddToCart(ctx *gin.Context) {
	req := new(models.CartRequest)
	if err := ctx.Bind(req); err != nil {
		helpers.DefaultErrorBadRequest(ctx, helpers.InvalidRequestBody)
		return
	}

	if err := services.AddToCart(req); err != nil {
		helpers.ReturnError(ctx, err)
		return
	}

	helpers.SuccessCreate(ctx, nil)
}

// AddToCart godoc
// @Summary Add product to cart
// @Description Add product to cart
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body models.CartRequest true "Cart request"
// @Success 201 {object} helpers.Response
// @Failure 400 {object} helpers.Response
// @Router /checkout [post]
func Checkout(ctx *gin.Context) {
	req := new(models.CartRequest)
	if err := ctx.Bind(req); err != nil {
		helpers.DefaultErrorBadRequest(ctx, helpers.InvalidRequestBody)
		return
	}

	if err := services.Checkout(req); err != nil {
		helpers.ReturnError(ctx, err)
		return
	}

	helpers.SuccessCreate(ctx, nil)
}
