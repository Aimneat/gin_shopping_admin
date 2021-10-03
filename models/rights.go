package models

import "github.com/jinzhu/gorm"

type Rights struct {
	gorm.Model
	AuthName string //权限说明
	Level    string //权限层级
	Pid      uint   //父权限id
	Path     string //对应访问路径
	Children []Rights
}
