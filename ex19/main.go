package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}


func main() {

	//gin.ForceConsoleColor()
	//gin.DisableConsoleColor()
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	defer db.Close()


	//Create table
	db.AutoMigrate(&UserInfo{})

	//
	u1 := UserInfo{1, "qimi", "nan", "wa yong"}
	u2 := UserInfo{2, "沙河娜扎", "女", "足球"}

	db.Create(&u1)
	db.Create(&u2)

	//get
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n", u)
	db.Last(&u)
	fmt.Printf("u:%#v\n", u)

	//
	db.Model(&u).Update("hobby", "双色球")


	//r := gin.Default()


	db.Delete(&u)
	//r.Run()
}
