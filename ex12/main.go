package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "index.html")

	// 一次请求对应一次响应
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	} )
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password :=c.PostForm("password")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Name": username,
		"Password": password,
	})
	})

	r.Run(":9999")
}
