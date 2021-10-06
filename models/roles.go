package models

import "github.com/jinzhu/gorm"

type Roles struct {
	gorm.Model
	RoleName string
	RoleDesc string //角色描述
}
