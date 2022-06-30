package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := router.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
