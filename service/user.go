package service

import (
	"gin-shop-admin/initialize"
	"gin-shop-admin/models"

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

func ExistUserByUsername(username string) (bool, error) {
	var user models.User
	err := initialize.Db.Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func CreateRootUser(user models.User) error {

	if err := initialize.Db.Create(&user).Error; err != nil {
		// app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return err
	}
	return nil
}
