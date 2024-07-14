package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	connectionString := os.Getenv("DSN")

	instance, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Database initial error: ", err.Error())
	}
	db = instance
	migrate(db)
}

func Db() *gorm.DB {
	return db
}
