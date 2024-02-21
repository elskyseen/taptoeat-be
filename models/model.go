package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateConnection(dbname string, user string, pass string) {
	conn := user + ":" + pass + "@tcp(127.0.0.1:3306)/" + dbname + "?parseTime=true"

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		fmt.Println("Ohh sorry can't connect into database")
	} else {
		fmt.Println("Database connected")
	}

	db.AutoMigrate(&User{})

	DB = db
}
