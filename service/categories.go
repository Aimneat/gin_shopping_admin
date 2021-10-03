package service

import (
	"gin-shop-admin/initialize"
	"gin-shop-admin/models"
	"gin-shop-admin/models/response"

	"github.com/jinzhu/gorm"
)

func CreateCategories() {

}

func UpdataCategories() {

}

func GetAllCategories() (categoriesList []response.Categories, err error) {
	var allCategories []models.Categories
	categoriesTreeMap := make(map[uint][]response.Categories)

	err = initialize.Db.Find(&allCategories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	for _, v := range allCategories {
		categoriesMap := response.Categories{
			ID:                v.ID,
			CategoriesName:    v.CategoriesName,
			CategoriesPid:     v.CategoriesPid,
			CatgoriesLevel:    v.CatgoriesLevel,
			CategoriesDeleted: false, //这里写死，本应需要判断
			Children:          nil,
		}
		categoriesTreeMap[v.CategoriesPid] = append(categoriesTreeMap[v.CategoriesPid], categoriesMap)
	}

	categoriesList = categoriesTreeMap[0]
	for i := 0; i < len(categoriesList); i++ {
		err = GetChildrenCategories(&categoriesList[i], categoriesTreeMap)
	}

	return categoriesList, err
}

func GetChildrenCategories(categoriesList *response.Categories, categoriesTreeMap map[uint][]response.Categories) (err error) {
	categoriesList.Children = categoriesTreeMap[categoriesList.ID]
	for i := 0; i < len(categoriesList.Children); i++ {
		err = GetChildrenCategories(&categoriesList.Children[i], categoriesTreeMap)
	}
	return err
}
