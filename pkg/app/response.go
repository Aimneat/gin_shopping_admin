package app

import (
	"gin-shop-admin/pkg/e"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpCode int, errCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"data": data,
		"meta": gin.H{
			"msg":    e.GetMsg(errCode),
			"status": errCode,
		},
	})
	return
}
