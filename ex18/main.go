package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHnadler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
})
}

func m1(c *gin.Context) {
	fmt.Println("IN m1 ... ")
}


func m2(c *gin.Context) {
	fmt.Println("IN m2 ... ")
}



func main() {
	r := gin.Default()

	r.GET("/index", m1, m2,  indexHnadler)

	r.Run(":9999")
}
