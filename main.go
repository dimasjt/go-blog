package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main() {
	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run()
}
