package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	const MySQL = "root:@tcp(127.0.0.1:3306)/sekolahCN?charset=utf8mb4&parseTime=True&loc=Local"
	DSN := MySQL

	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Connected to Database")
}
