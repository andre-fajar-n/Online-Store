package route

import (
	"log"
	"net/http"

	"github.com/andre-fajar-n/Online-Store/controllers"
	_ "github.com/andre-fajar-n/Online-Store/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Route(router *gin.Engine) {
	log.Println("Initialize Routing...")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1Router := router.Group("/api/v1/online-store")

	v1Router.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Online Store")
	})

	v1Router.POST("cart", controllers.AddToCart)
	v1Router.POST("checkout", controllers.Checkout)
}
