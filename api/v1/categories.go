package v1

import (
	"gin-shop-admin/pkg/app"
	"gin-shop-admin/pkg/e"
	"gin-shop-admin/service"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 获取商品分类列表
// @Produce  json
// @Param type query int false "显示几层分类列表" Enums(1,2,3) default(3)
// @Param pagenum query int "当前页码值"
// @Param pagesize query int "每页显示数据条数"
// @Router /api/private/v1/categories [get]
func Categories(c *gin.Context) {
	// level := c.DefaultQuery("type", "3")
	pagenum := c.DefaultQuery("pagenum", "1")
	pagesize := c.DefaultQuery("pagesize", "5")
	var offsetNum = 0
	//to do: 需要验证数据的正确性

	pnum, err := strconv.Atoi(pagenum)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_CATEGORIES_FAIL, nil)
		return
	}
	size, err := strconv.Atoi(pagesize)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_CATEGORIES_FAIL, nil)
		return
	}
	if pnum > 0 {
		offsetNum = (pnum - 1) * size
	}

	categoriesList, err := service.GetAllCategories()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_CATEGORIES_FAIL, nil)
		return
	}

	listLength := math.Min(float64(len(categoriesList)), float64(offsetNum+size))
	data := gin.H{
		"result": categoriesList[offsetNum:int(listLength):int(listLength)],
		"total":  len(categoriesList),
	}
	app.Response(c, http.StatusOK, e.SUCCESS,
		data,
	)
}

func AddCategories(c *gin.Context) {

}
