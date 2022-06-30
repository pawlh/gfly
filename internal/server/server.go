package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	registerRoutes(router)

	router.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))

	err := router.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
