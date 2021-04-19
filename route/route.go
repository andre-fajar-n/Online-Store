package route

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	log.Println("Initialize Routing...")

	v1URI := router.Group("/api/v1/online-store")

	v1URI.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PING")
	})
}
