package main

import "github.com/gin-gonic/gin"

func main() {

	gin.ForceConsoleColor()
	//gin.DisableConsoleColor()


	r := gin.Default()

	r.Run()
}
