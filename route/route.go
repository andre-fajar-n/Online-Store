package route

import (
	"log"
	"net/http"

	"github.com/andre-fajar-n/Online-Store/controllers"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	log.Println("Initialize Routing...")

	v1Router := router.Group("/api/v1/online-store")

	v1Router.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PING")
	})

	cartController := controllers.Controller{
		StatusCode: http.StatusOK,
	}
	v1Router.POST("cart", cartController.AddToCart)
}
