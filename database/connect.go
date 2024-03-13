package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func connectToDB(){
	var err error
	dsn := 
	DB, err = gorm.Open()

}
