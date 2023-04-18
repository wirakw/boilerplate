package database

import (
	"boilerplate/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func init() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/boilerplate?charset=utf8mb4&parseTime=True&loc=Local"
	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Conn.AutoMigrate(&entities.Activity{}, &entities.Todo{})
}
