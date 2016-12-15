package main

func initializeRoutes() {
	router.GET("/", showIndexPage)

	router.GET("/new", newArticle)

	router.GET("/articles/:article_id", showArticle)

	router.POST("/create", createArticle)
}
