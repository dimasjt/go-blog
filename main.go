package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run()
}

func render(c *gin.Context, templateName string, data gin.H) {
	c.HTML(http.StatusOK, templateName, data)
}
