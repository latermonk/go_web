package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		data := gin.H{"name":"small wang zhi", "message":"Hello world !", "age": 182}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9999")
}
