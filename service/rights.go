package service

import (
	"gin-shop-admin/initialize"
	"gin-shop-admin/models"
	"gin-shop-admin/models/response"
	"strconv"
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
