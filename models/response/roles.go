package response

import "gin-shop-admin/models"

type Roles struct {
	ID       uint            `json:"id"`
	RoleName string          `json:"roleName"`
	RoleDesc string          `json:"roleDesc"`
	Rights   []models.Rights `json:"children" gorm:"many2many:rights_roles;"`
}

type RRoles struct {
	ID       uint    `json:"id"`
	RoleName string  `json:"roleName"`
	RoleDesc string  `json:"roleDesc"`
	Rights   []Menus `json:"children"`
}
