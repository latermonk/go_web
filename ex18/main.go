package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func indexHnadler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
})
}

func m1(c *gin.Context) {
	fmt.Println("IN m1 ... ")
	start := time.Now()
	c.Set("name", "xiao wang zhi")
	c.Next()

	cost := time.Since(start)
	log.Println(cost)
}


func m2(c *gin.Context) {
	fmt.Println("IN m2 ... ")
}



func main() {
	r := gin.Default()

	r.GET("/index", m1, m2,  indexHnadler)

	r.Run(":9999")
}
