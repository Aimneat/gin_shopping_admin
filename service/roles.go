package service

import (
	"gin-shop-admin/initialize"
	"gin-shop-admin/models"
	"gin-shop-admin/models/response"

	"github.com/jinzhu/gorm"
)

func GetAllRoles() (roles []response.Roles, err error) {
	err = initialize.Db.Table("roles").Select("roles.id,roles.role_name,roles.role_desc").Preload("Rights", func(query *gorm.DB) *gorm.DB {
		return query.Select("*")
	}).Where("roles.deleted_at is null").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func GetRoleByRoleName(roleName string) (role models.Roles, exist bool, err error) {
	err = initialize.Db.Table("roles").Where("role_name = ?", roleName).Preload("Rights", func(query *gorm.DB) *gorm.DB {
		return query.Select("*")
	}).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return role, false, err
	}
	if role.ID > 0 {
		return role, true, nil
	}
	return role, false, nil
}

func GetRoleByID(roleID uint) (role models.Roles, exist bool, err error) {
	err = initialize.Db.Table("roles").Where("id = ?", roleID).Preload("Rights", func(query *gorm.DB) *gorm.DB {
		return query.Select("*")
	}).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return role, false, err
	}
	if role.ID > 0 {
		return role, true, nil
	}
	return role, false, nil
}

func CreateRoles(newRole models.Roles) error {
	if err := initialize.Db.Create(&newRole).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRole(newRole models.Roles) error {
	if err := initialize.Db.Save(&newRole).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRole(roleID uint) error {
	var role models.Roles
	err := initialize.Db.Where("id = ?", roleID).Delete(&role).Error
	if err != nil {
		return err
	}
	return nil
}
