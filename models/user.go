package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	RolesID  uint
	Username string
	Password string
	Mobile   string
	Email    string
	IsActive bool
	Type     uint
	MgState  bool
}
