package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var dsn = "root:newpassword@tcp(172.17.0.2:3306)/movies?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}
