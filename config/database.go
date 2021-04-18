package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func ConnectDB() *gorm.DB {
	log.Println("Connecting DB...")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	envPath := ".env"
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file")
	}

	username := os.Getenv("USER_DATABASE")
	password := os.Getenv("PASS_DATABASE")
	host := os.Getenv("HOST_DATABASE")
	port := os.Getenv("PORT_DATABASE")
	name := os.Getenv("NAME_DATABASE")

	conn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, errConn := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newLogger,
	})

	if errConn != nil {
		panic("Failed to connect to database!")
	}
	log.Println("Database connection success!")

	return db
}
