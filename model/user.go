package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	CreatedBy uint    `json:"created_by" gorm:"index"` // 1-to-1 with User
	Owners    []User  `json:"owners" gorm:"many2many:user_goods;"`
}

type User struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Wallet          string    `json:"wallet"`
	PrivateKey      string    `json:"private_key"`
	CreatedProducts []Product `json:"created_products" gorm:"foreignKey:UserID"`
	BoughtProducts  []Product `json:"bought_products" gorm:"foreignKey:UserID"`
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Product{})
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func CreateGood(db *gorm.DB, good *Product) error {
	return db.Create(good).Error
}

func GetGoodsByUserID(db *gorm.DB, userID uint) ([]Product, error) {
	var goods []Product
	if err := db.Where("created_by = ?", userID).Find(&goods).Error; err != nil {
		return nil, err
	}
	return goods, nil
}
