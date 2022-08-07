package v1

import (
	"gin-shop-admin/middleware"
	"gin-shop-admin/models"
	"gin-shop-admin/models/request"
	"gin-shop-admin/pkg/app"
	"gin-shop-admin/pkg/e"
	"gin-shop-admin/pkg/util"
	"gin-shop-admin/service"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type IUser interface {
// 	Login(c *gin.Context)
// 	LoginOut(c *gin.Context)
// 	Register(c *gin.Context)
// 	CheckToken(c *gin.Context)
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

	_, exist, err := service.ExistUserByUsername(rUser.Username)
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

func Users(c *gin.Context) {
	query := c.DefaultQuery("query", "") //字符串搜索
	pagenum := c.DefaultQuery("pagenum", "1")
	pagesize := c.DefaultQuery("pagesize", "5")
	var offsetNum = 0
	//to do: 需要验证数据的正确性

	pnum, err := strconv.Atoi(pagenum)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USERS_FAIL, nil)
		return
	}
	size, err := strconv.Atoi(pagesize)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USERS_FAIL, nil)
		return
	}
	if pnum > 0 {
		offsetNum = (pnum - 1) * size
	}

	usersList, err := service.GetAllUsers(query)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USERS_FAIL, nil)
		return
	}

	listLength := math.Min(float64(len(usersList)), float64(offsetNum+size))
	data := gin.H{
		"users": usersList[offsetNum:int(listLength):int(listLength)],
		"totle": len(usersList),
	}
	app.Response(c, http.StatusOK, e.SUCCESS,
		data,
	)
}

func AddUser(c *gin.Context) {
	var ruser request.AddUser
	err := c.ShouldBind(&ruser)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	//todo ：需要验证数据正确性

	_, exist, err := service.ExistUserByUsername(ruser.Username)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}
	if exist == true {
		app.Response(c, http.StatusOK, e.ERROR_ADD_EXIST_USER, nil)
		return
	}

	password, err := util.SetPassword(ruser.Username)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}
	user := models.User{
		Username: ruser.Username,
		Password: password,
		Email:    ruser.Email,
		Mobile:   ruser.Mobile,
		RolesID:  4, //默认为游客
	}

	err = service.CreateUser(user)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}

	user, exist, err = service.ExistUserByUsername(ruser.Username)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	if exist == false {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_USER_FAIL, nil)
		return
	}
	app.Response(c, http.StatusCreated, e.CREATED_SUCCESS, gin.H{
		"id":          user.ID,
		"username":    user.Username,
		"mobile":      user.Mobile,
		"type":        user.Type,
		"openid":      "",
		"email":       user.Email,
		"create_time": user.CreatedAt,
		"modify_time": user.UpdatedAt,
		"is_delete":   user.DeletedAt != nil,
		"is_active":   user.IsActive,
	})

}

func UserStateChanged(c *gin.Context) {
	uid := c.Param("uId")
	mgState := c.Param("type")

	//todo ：需要验证数据正确性

	id, err := strconv.Atoi(uid)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPDATE_USER_FAIL, nil)
		return
	}

	user, exist, err := service.GetUserById(uint(id))
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	if exist == false {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	user.MgState, err = strconv.ParseBool(mgState)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPDATE_USER_FAIL, nil)
		return
	}

	err = service.UpdateUser(user)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPDATE_USER_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, gin.H{
		"id":       user.ID,
		"rid":      user.RolesID,
		"username": user.Username,
		"mobile":   user.Mobile,
		"email":    user.Email,
		"mg_state": user.MgState,
	})
}

func EditUser(c *gin.Context) {
	uid := c.Param("id")
	var ruser request.EditUser
	err := c.ShouldBind(&ruser)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPDATE_USER_FAIL, nil)
		return
	}

	//todo ：需要验证数据正确性

	id, err := strconv.Atoi(uid)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPDATE_USER_FAIL, nil)
		return
	}

	user, exist, err := service.GetUserById(uint(id))
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	if exist == false {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	user.Mobile = ruser.Mobile
	user.Email = ruser.Email

	err = service.UpdateUser(user)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPDATE_USER_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, gin.H{
		"id":       user.ID,
		"rid":      user.RolesID,
		"username": user.Username,
		"mobile":   user.Mobile,
		"email":    user.Email,
	})
}

func GetUserById(c *gin.Context) {
	uid := c.Param("id")

	//todo ：需要验证数据正确性

	id, err := strconv.Atoi(uid)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPDATE_USER_FAIL, nil)
		return
	}

	user, exist, err := service.GetUserById(uint(id))
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	if exist == false {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, gin.H{
		"id":       user.ID,
		"rid":      user.RolesID,
		"username": user.Username,
		"mobile":   user.Mobile,
		"email":    user.Email,
	})
}

func DeleteUser(c *gin.Context) {
	uid := c.Param("id")

	//todo ：需要验证数据正确性

	id, err := strconv.Atoi(uid)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_USER_FAIL, nil)
		return
	}

	_, exist, err := service.GetUserById(uint(id))
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}
	if exist == false {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	err = service.DeleteUserById(uint(id))
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_USER_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
