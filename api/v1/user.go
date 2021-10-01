package v1

import (
	"gin-shop-admin/middleware"
	"gin-shop-admin/models"
	"gin-shop-admin/models/request"
	"gin-shop-admin/pkg/app"
	"gin-shop-admin/pkg/e"
	"gin-shop-admin/pkg/util"
	"gin-shop-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type IUser interface {
// 	Login(c *gin.Context)
// 	LoginOut(c *gin.Context)
// 	Register(c *gin.Context)

// 	CheckToken(c *gin.Context)

// 	MyInformation(c *gin.Context)
// }

// type UserManager struct {
// 	DB *gorm.DB
// }

// CheckToken 用户详情
func CheckToken(c *gin.Context) {
	app.Response(c,
		http.StatusOK,
		e.SUCCESS,
		nil,
	)
}

// @Summary 用户登录
// @Produce  json
// @Router /api/private/v1/login [post]
func Login(c *gin.Context) {
	var rUser request.Login

	err := c.ShouldBindJSON(&rUser)
	if err != nil {
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	exist, err := service.ExistUser(rUser.Username, util.EncodeMD5(rUser.Password))
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	if !exist {
		app.Response(c, http.StatusUnprocessableEntity, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	j := middleware.NewJWT()
	token, err := j.GenerateToken(models.User{Username: rUser.Username})
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GENERATE_TOKEN_FAIL, nil)
		return
	}
	app.Response(c, http.StatusOK, e.SUCCESS, gin.H{
		"id":       500,
		"rid":      0,
		"username": rUser.Username,
		"mobile":   "",
		"email":    "",
		"token":    token,
	})

}

func LoginOut(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func Register(c *gin.Context) {
	var rUser request.Register
	err := c.ShouldBindJSON(&rUser)
	if err != nil {
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if rUser.InvitationCode != "admin" { //邀请码不应该为固定字符串
		app.Response(c, http.StatusUnprocessableEntity, e.ERROR_USER_INVITATIONCODE, nil)
		return
	}

	exist, err := service.ExistUserByUsername(rUser.Username)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	if exist {
		app.Response(c, http.StatusUnprocessableEntity, e.ERROR_ADD_EXIST_USER, nil)
		return
	}

	user := models.User{
		Username: rUser.Username,
		Password: util.EncodeMD5(rUser.Password),
	}
	if err := service.CreateRootUser(user); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	j := middleware.NewJWT()
	token, err := j.GenerateToken(models.User{Username: user.Username})
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GENERATE_TOKEN_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, gin.H{
		"id":       500,
		"rid":      0,
		"username": rUser.Username,
		"mobile":   "",
		"email":    "",
		"token":    token,
	})
}

// func (u UserManager) MyInformation(c *gin.Context) {
// 	id := c.Param("id")

// 	var user models.User

// 	err := u.DB.Preload("Orders").Where("id = ?", id).First(&user).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		// app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
// 		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
// 		return
// 	}

// 	app.Response(c, http.StatusOK, e.SUCCESS, gin.H{"user": user})
// }

// func NewUserManager() IUser {
// 	db := models.GetDB()
// 	db.AutoMigrate(&models.User{})
// 	return UserManager{DB: db}
// }
