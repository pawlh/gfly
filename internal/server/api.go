package server

import (
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	router.GET("/api/points", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, getPoints())
	})
}
