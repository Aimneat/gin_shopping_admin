package models

import "github.com/jinzhu/gorm"

type Rights struct {
	gorm.Model
	authName string //权限说明
	level    string //权限层级
	pid      uint   //父权限id
	path     string //对应访问路径
}
