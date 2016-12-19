package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func registerUser(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if _, err := createUser(name, email, password); err == nil {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
