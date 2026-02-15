package main

import (
	"log"

	"github.com/adhitamafikri/cozy-prop-tech/backend/api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	log.Printf("Starting API on %s", cfg.Port)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "index"})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})

	address := "0.0.0.0:" + cfg.Port
	router.Run(address)
}
