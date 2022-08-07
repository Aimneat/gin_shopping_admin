package service

import (
	"gin-shop-admin/initialize"
	"gin-shop-admin/models"
	"gin-shop-admin/models/response"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetAllRights() (menus []response.Menus, err error) {
	var allRights []models.Rights
	menusTreeMap := make(map[string][]response.Menus)

	err = initialize.Db.Find(&allRights).Error
	if err != nil {
		return nil, err
	}
	for _, v := range allRights {
		ParentsId := strconv.Itoa(int(v.Pid))
		menusMap := response.Menus{
			ID:       v.ID,
			AuthName: v.AuthName,
			Path:     v.Path,
			Children: nil,
		}
		menusTreeMap[ParentsId] = append(menusTreeMap[ParentsId], menusMap)
	}

	menus = menusTreeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = GetChildrenRights(&menus[i], menusTreeMap)
	}

	return menus, err
}

func GetChildrenRights(menus *response.Menus, menusTreeMap map[string][]response.Menus) (err error) {
	menus.Children = menusTreeMap[strconv.Itoa(int(menus.ID))]
	for i := 0; i < len(menus.Children); i++ {
		err = GetChildrenRights(&menus.Children[i], menusTreeMap)
	}
	return err
}

func GetAllRightsList() (rightsList []models.Rights, err error) {
	err = initialize.Db.Find(&rightsList).Error
	if err != nil {
		return nil, err
	}
	return rightsList, nil
}

func GetRightsById(rightsId uint) (menus response.Menus, err error) {
	err = initialize.Db.Table("rights").Where("id = ?", rightsId).First(&menus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return menus, err
	}
	childrenMenus, err := GetRightsByPid(rightsId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return menus, err
	}
	menus.Children = childrenMenus
	return menus, nil
}

func GetRightsByPid(pid uint) (menus []response.Menus, err error) {
	err = initialize.Db.Table("rights").Where("pid = ?", pid).Find(&menus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return menus, err
	}
	for _, v := range menus {
		childrenMenus, err := GetRightsByPid(v.ID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return menus, err
		}
		v.Children = childrenMenus
	}
	return menus, nil
}
