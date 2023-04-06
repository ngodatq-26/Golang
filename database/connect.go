package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"yogurt-project/models"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = Connect()
	return Db
}

const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "yogurt-project"
const DB_HOST = "localhost"
const DB_PORT = "3306"

func Connect() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error Connecting database")
		return nil
	}

	db.AutoMigrate(&models.User{})

	return db
}
