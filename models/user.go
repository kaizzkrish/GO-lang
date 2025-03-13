package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`
	Username    string `gorm:"unique"`
	Email       string `gorm:"unique"`
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
