package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "m_user:M_pwd123@@tcp(127.0.0.1:3306)/movies_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB connection failed:", err)
		return
	}

	DB = db
	fmt.Println("Database connected with GORM ✅")
}