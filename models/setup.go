package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3307)/isengin"))
	if err != nil {
		fmt.Println("Gagal ke database")
	}
	db.AutoMigrate(&User{})
	DB = db
}
