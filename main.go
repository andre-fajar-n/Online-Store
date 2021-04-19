package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andre-fajar-n/Online-Store/config"
	"github.com/andre-fajar-n/Online-Store/models"
	"github.com/andre-fajar-n/Online-Store/route"
	"github.com/andre-fajar-n/Online-Store/seed"
	"github.com/gin-gonic/gin"
)

const (
	serve   = "serve"
	migrate = "migrate"
)

var commandType = map[string]string{
	serve:   "Run Server",
	migrate: "Migrate Database",
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Invalid command")
		fmt.Println("Please add one of this command type in command, e.g go run main go", serve)

		i := 1
		for key, value := range commandType {
			fmt.Println(i, key, ": used to", value)
			i++
		}
		os.Exit(0)
	}

	parseCommand(args)
}

func parseCommand(text []string) {
	log.Println("Process", text)

	switch text[1] {
	case "serve":
		runServer()
	case "migrate":
		runMigrate()
	case "seed":
		runSeed()
	}
}

func runServer() {
	r := gin.Default()

	// connect db
	config.ConnectDB()

	// initialize route
	route.Route(r)

	// route not found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"en": "route not found",
			"id": "route tidak ditemukan",
		})
	})

	r.Run(":7000")
}

func runMigrate() {
	db := config.ConnectDB()

	log.Println("Start Migration...")

	if err := db.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.Order{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.OrderDetail{}); err != nil {
		panic(err)
	}

	log.Println("Success Migration!")
}

func runSeed() {
	db := config.ConnectDB()
	seed.UserSeed(db)
	seed.ProductSeed(db)
}
