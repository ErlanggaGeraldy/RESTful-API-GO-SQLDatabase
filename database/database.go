package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/test-interview?parseTime=true"))
	if err != nil {
		panic("Cant Connect to Database")
	}
	fmt.Println("Connected to Database")
	DB = database
}
