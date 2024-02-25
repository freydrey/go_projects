package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//we are using gorm orm for connecting to mysql

var(
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "bkstore:bkstore@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}