package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//create database
	//connect database
	v1Group.POST("/todo", func(c *gin.Context) {

	})


	r := gin.Default()

	r.Static("/static", "static")

	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//v1
	v1Group := r.Group("v1")
	{

	}


	r.Run(":9999")
}
