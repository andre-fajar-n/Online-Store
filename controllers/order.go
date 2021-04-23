package controllers

import (
	"github.com/andre-fajar-n/Online-Store/helpers"
	"github.com/andre-fajar-n/Online-Store/models"
	"github.com/andre-fajar-n/Online-Store/services"
	"github.com/gin-gonic/gin"
)

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
