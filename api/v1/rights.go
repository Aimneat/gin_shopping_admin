package v1

import (
	"gin-shop-admin/pkg/app"
	"gin-shop-admin/pkg/e"
	"gin-shop-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Menus(c *gin.Context) {

	menus, err := service.GetAllRights()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_MENUS_ERROR, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS,
		menus,
	)
}
