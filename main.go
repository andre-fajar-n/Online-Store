package main

import (
	"net/http"

	"github.com/andre-fajar-n/Online-Store/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()

	r.Run(":7000")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// ping test
	r.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ping success")
	})

	// connect db
	config.ConnectDB()

	return r
}
