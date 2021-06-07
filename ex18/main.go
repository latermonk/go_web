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

	fmt.Println("Out m1 ...")
}


func m2(c *gin.Context) {
	fmt.Println("IN m2 ... ")
	c.Next()
	fmt.Println("Out m2 ...")
}



func main() {
	r := gin.Default()

	//global register middleware
	r.Use(m1, m2)

	//r.GET("/index", m1, m2,  indexHnadler)
	r.GET("/index",  indexHnadler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "shop",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user",
		})
	})

	r.Run(":9999")
}
