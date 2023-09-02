package model

import (
	"errors"

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

type UserGoodDB struct {
	ProductID uint `gorm:"primary_key"`
	UserID    uint `gorm:"primary_key"`
}

func InsertUserGood(db *gorm.DB, userID, productID uint) error {
	userGood := UserGoodDB{
		ProductID: productID,
		UserID:    userID,
	}

	return db.Table("user_goods").Create(&userGood).Error
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

func GetGoodsByID(db *gorm.DB, productID uint) (Product, error) {
	var good Product
	if err := db.Where("id = ?", productID).First(&good).Error; err != nil {
		return Product{}, err
	}
	return good, nil
}

type ProductName struct {
	ID   uint   `json:"id"`
	Name string `json:"product_name"`
}

func (ProductName) TableName() string {
	return "products"
}

func GetProductsCountByUserID(db *gorm.DB, userID uint) ([]ProductName, error) {
	var products []Product
	if err := db.Where("created_by = ?", userID).Find(&products).Error; err != nil {
		return nil, err
	}

	var responses []ProductName
	for _, p := range products {
		responses = append(responses, ProductName{
			ID:   p.ID,
			Name: p.Name,
		})
	}

	return responses, nil
}

type ProductSoldCount struct {
	ProductName
	Count int `json:"count"`
}

type ProductCount struct {
	ProductId uint   `json:"product_id"`
	Name      string `json:"product_name"`
	Count     int    `json:"count"`
}

func GetProductsSoldCountUserID(db *gorm.DB, userID uint) ([]ProductSoldCount, error) {
	var products []Product
	var responses []ProductSoldCount

	// Fetch products associated with user
	err := db.Where("created_by = ?", userID).Find(&products).Error
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		var count int
		err := db.Table("user_goods").Where("user_id <> ?", userID).Where("product_id = ?", product.ID).Count(&count).Error
		if err != nil {
			return nil, err
		}

		if count > 0 {
			responses = append(responses, ProductSoldCount{
				ProductName: ProductName{
					ID:   product.ID,
					Name: product.Name},
				Count: count,
			})
		}
	}

	return responses, nil
}

type UserGood struct {
	ProductID uint
	UserID    uint
}

func GetProductsBoughtByUserID(db *gorm.DB, userID uint) ([]ProductName, error) {
	var responses []ProductName
	var userGoods []UserGood

	// Fetch all product IDs associated with the user
	if err := db.Where("user_id = ?", userID).Find(&userGoods).Error; err != nil {
		return nil, err
	}

	// Extract all product IDs
	productIDs := make([]uint, len(userGoods))
	for i, ug := range userGoods {
		productIDs[i] = ug.ProductID
	}

	// Fetch product names using the extracted product IDs
	if err := db.Where("created_by <> ?", userID).Where("id IN (?)", productIDs).Find(&responses).Error; err != nil {
		return nil, err
	}

	return responses, nil

}

func GetProductsBoughtCountByProductID(db *gorm.DB, productID uint) (ProductSoldCount, error) {
	var countResult ProductSoldCount
	countResult.ID = productID

	var productCnt ProductCount
	if err := db.Table("user_goods").Select("product_id, count(*) as count").Where("product_id = ?", productID).Group("product_id").First(&productCnt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If no record is found, set count to 0
			countResult.Count = 0
		}
	}
	countResult.Count = productCnt.Count

	var products []Product
	// Fetch product names using the extracted product IDs
	if err := db.Where("id = ?", productID).Find(&products).Error; err != nil {
		return countResult, err
	}
	if len(products) == 1 {
		countResult.Name = products[0].Name
	}

	return countResult, nil
}
