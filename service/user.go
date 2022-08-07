package service

import (
	"gin-shop-admin/initialize"
	"gin-shop-admin/models"
	"gin-shop-admin/models/response"

	"github.com/jinzhu/gorm"
)

func ExistUser(username, password string) (bool, error) {
	var user models.User
	err := initialize.Db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func ExistUserByUsername(username string) (user models.User, existe bool, err error) {

	err = initialize.Db.Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, false, err
	}
	if user.ID > 0 {
		return user, true, nil
	}
	return user, false, nil
}

func GetUserById(id uint) (user models.User, exist bool, err error) {
	err = initialize.Db.Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, false, err
	}
	if user.ID > 0 {
		return user, true, nil
	}
	return user, false, nil
}

func CreateRootUser(user models.User) error {

	if err := initialize.Db.Create(&user).Error; err != nil {
		// app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return err
	}
	return nil
}

func GetAllUsers(query string) (usersList []response.Users, err error) {
	var users []models.User

	// err = initialize.Db.Model(&users).Select("users.id,users.username,users.mobile,users.type,users.email,users.created_at,users.mg_state,roles.role_name").
	// 	Joins("join roles on roles.id=users.roles_id where users.deleted_at is null and roles.deleted_at is null").Find(&usersList).Error
	err = initialize.Db.Model(&users).Select("users.id,users.username,users.mobile,users.type,users.email,users.created_at,users.mg_state,roles.role_name").
		Joins("join roles on roles.id=users.roles_id").
		Where("users.deleted_at is null and roles.deleted_at is null and users.username like ?", "%"+query+"%").Find(&usersList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return usersList, err
}

func CreateUser(user models.User) error {
	if err := initialize.Db.Create(&user).Error; err != nil {
		// app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return err
	}
	return nil
}

func UpdateUser(user models.User) error {
	if err := initialize.Db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserById(id uint) (err error) {
	var user models.User
	err = initialize.Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
