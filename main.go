package main

import (
	"net/http"
	"os"

	"gopkg.in/gin-gonic/gin.v1"
)

var router *gin.Engine

func main() {
	if os.Getenv("GIN_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run()
}

func render(c *gin.Context, templateName string, data gin.H) {
	c.HTML(http.StatusOK, templateName, data)
}
