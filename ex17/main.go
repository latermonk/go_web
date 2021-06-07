package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",

		})
	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method" : "PUT",
		})

	})

	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"message":"GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"message": "POST"})

		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ANY",

		})
	})


	//r.NoRoute("", func(context *gin.Context) {
	//
	//})

	r.Run(":9999")
}
