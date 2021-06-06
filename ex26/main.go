package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	DB *gorm.DB
)
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func  initMySQL()(err error){
	dsn := "root:root1234@tcp(127.0.0.1:13306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)

	return
	//return DB.DB().Ping()
}


func main() {
	//create data
	//docker run --name mysql8019 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -d mysql:8.0.19

	//connect database
	err := initMySQL()
	if err != nil{
		panic(err)
	}
	defer DB.Close()

	//
	DB.AutoMigrate(&Todo{})



	r := gin.Default()

	r.Run(":9999")
}
