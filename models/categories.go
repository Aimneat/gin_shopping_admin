package models

import "github.com/jinzhu/gorm"

type Categories struct {
	gorm.Model
	CategoriesName string
	CategoriesPid  uint //分类父ID
	CatgoriesLevel uint //分类当前层级
	Children       []Categories
}
