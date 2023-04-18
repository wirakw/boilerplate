package database

import (
	"boilerplate/entities"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func init() {
	var err error
	dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")/" + os.Getenv("MYSQL_DBNAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Conn.AutoMigrate(&entities.Activity{}, &entities.Todo{})
}
