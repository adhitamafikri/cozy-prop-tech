package config

import "github.com/gin-gonic/gin"

func NewGinEngine() *gin.Engine {
	router := gin.Default()
	return router
}
