package main

import (
	"net/http"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	render(c, "index.html", gin.H{
		"title":   "Home Page",
		"payload": articles,
	})
}

func showArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			render(c, "article.html",
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewArticle(title, content); err == nil {
		articleID := strconv.Itoa(a.ID)
		c.Redirect(http.StatusMovedPermanently, "/articles/"+articleID)
	}
}
