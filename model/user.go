package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Wallet     string `json:"wallet"`
	PrivateKey string `json:"private_key"`
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
