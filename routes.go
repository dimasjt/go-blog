package main

func initializeRoutes() {
	router.GET("/", showIndexPage)

	router.GET("/articles/:article_id", showArticle)
}
