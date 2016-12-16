package main

import "github.com/gin-gonic/gin"

func initializeRoutes() {
	router.Static("uploads/", "./public/uploads")

	router.GET("/", showIndexPage)

	router.GET("/new", func(c *gin.Context) {
		render(c, "new_article.html", gin.H{
			"title": "New Article",
		})
	})

	router.GET("/articles/:article_id", showArticle)

	router.POST("/create", createArticle)

	router.GET("/register", func(c *gin.Context) {
		render(c, "register.html", gin.H{
			"title": "Register",
		})
	})

	router.POST("/register", registerUser)
}
