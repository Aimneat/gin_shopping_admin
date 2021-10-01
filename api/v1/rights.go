package v1

import (
	"gin-shop-admin/models/response"
	"gin-shop-admin/pkg/app"
	"gin-shop-admin/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Menus(c *gin.Context) {
	menus := response.Menus{
		ID:       101,
		AuthName: "商品列表",
		Path:     "",
		Children: nil,
	}
	app.Response(c, http.StatusOK, e.SUCCESS, gin.H{
		"id":       101,
		"authName": "商品管理",
		"path":     nil,
		"children": menus,
	})
	// app.Response(c, http.StatusOK, e.SUCCESS,
	// 	menus,
	// )
}
