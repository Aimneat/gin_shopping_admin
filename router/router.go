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
		apiv1.POST("/categories", v1.AddCategories)

		apiv1.GET("/users", v1.Users)
		apiv1.GET("/users/:id", v1.GetUserById)
		apiv1.PUT("/users/:id", v1.EditUser)
		apiv1.PUT("/user/:uId/state/:type", v1.UserStateChanged)
		apiv1.POST("/users", v1.AddUser)
		apiv1.DELETE("/users/:id", v1.DeleteUser)

		apiv1.GET("/roles", v1.GetAllRoles)
		apiv1.GET("/roles/:id", v1.GetRoleByID)
		apiv1.PUT("/roles/:id", v1.EditRole)
		apiv1.DELETE("/roles/:id", v1.DeleteRole)
		apiv1.POST("/roles", v1.AddRoles)
		apiv1.POST("/roles/:roleId/rights", v1.AllotRights)
		apiv1.DELETE("/role/:roleId/rights/:rightId", v1.RemoveRight)

		apiv1.GET("/rights/:type", v1.GetRights)
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
