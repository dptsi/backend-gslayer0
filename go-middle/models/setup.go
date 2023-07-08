package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:my-secret-pw@tcp(localhost:3306)/go_middle"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	DB = db
}
