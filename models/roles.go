package models

import "time"

type Roles struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
	RoleName  string     `json:"roleName"`
	RoleDesc  string     `json:"roleDesc"` //角色描述
	Rights    []Rights   `json:"children" gorm:"many2many:rights_roles;"`
}
