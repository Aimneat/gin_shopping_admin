package v1

import (
	"gin-shop-admin/models/response"
	"gin-shop-admin/pkg/app"
	"gin-shop-admin/pkg/e"
	"gin-shop-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Menus(c *gin.Context) {

	menus, err := service.GetAllRights()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_MENUS_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS,
		menus,
	)
}

func GetRights(c *gin.Context) {
	rightsType := c.Param("type")

	if rightsType == "list" { //返回的数据可能有误,可能会和tree的实现方式一样
		rightsList, err := service.GetAllRightsList()
		if err != nil {
			app.Response(c, http.StatusInternalServerError, e.ERROR_GET_RIGHTSLIST_FAIL, nil)
			return
		}

		List := make([]response.RightsList, len(rightsList), len(rightsList))
		for k, v := range rightsList {
			List[k].ID = v.ID
			List[k].AuthName = v.AuthName
			List[k].Level = v.Level
			List[k].Pid = v.Pid
			List[k].Path = v.Path
		}
		app.Response(c, http.StatusOK, e.SUCCESS, List)
	}
	if rightsType == "tree" {
		menus, err := service.GetAllRights()
		if err != nil {
			app.Response(c, http.StatusInternalServerError, e.ERROR_GET_MENUS_FAIL, nil)
			return
		}

		app.Response(c, http.StatusOK, e.SUCCESS,
			menus,
		)
	}
}
