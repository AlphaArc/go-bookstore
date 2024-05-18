package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	dsn := "root:db@p05t!@tcp(127.0.0.1:3306)/go-bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	// d, err := gorm.Open("mysql","root:db@p05t!/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql",dsn)
	if err != nil{
		panic(err)
	}
	fmt.Println("DB : go-bookstore Connection Established ")
	db = d
}

// Function to return db
func GetDB() *gorm.DB{
	return db
}
