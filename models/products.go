package models

import "github.com/jinzhu/gorm"

type Products struct {
	gorm.Model
	ProductName   string
	ProductPrice  uint
	ProductNumber uint
	ProductWeight uint
	ProductState  uint
	Sold          uint
	isHotSale     bool
}
