package database

import (
	users "go-api-template/app/modules/users/domain"

	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
}
