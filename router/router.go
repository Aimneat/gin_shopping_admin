package router

import (
	v1 "gin-shop-admin/api/v1"
	_ "gin-shop-admin/docs"
	"gin-shop-admin/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.CORSMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //shell:swag init 后生成swagger文件

	apiv1 := r.Group("/api/private/v1")
	apiv1.POST("/login", v1.Login)
	apiv1.POST("/register", v1.Register)

	apiv1.Use(middleware.JWTAuth())
	{
		apiv1.GET("/menus", v1.Menus)
		apiv1.GET("/categories", v1.Categories)
		//验证token
		// apiv1.GET("/ping", user.CheckToken)

		// apiv1.GET("/user/:id", user.MyInformation)

		// product := v1.NewProductManager()
		// apiv1.POST("/product", product.Create)
		// apiv1.GET("/product/:id", product.ShowByKey)
		// apiv1.GET("/products", product.ShowAll)
		// apiv1.DELETE("/product/:id", product.Delete)
		// apiv1.PUT("/product/:id", product.Update)

		// order := v1.NewOrderManager()
		// apiv1.POST("/order", order.Create)
		// apiv1.GET("/order/:id", order.ShowByKey)
		// apiv1.GET("/orders", order.ShowAll)
		// apiv1.DELETE("/order/:id", order.Delete)
		// apiv1.PUT("/order/:id", order.Updata)
	}

	return r
}
