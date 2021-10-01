package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Rid      uint
	Username string
	Password string
	Mobile   string
	Email    string
	IsDelete bool
	IsActive bool
	// Orders    []Order
}
