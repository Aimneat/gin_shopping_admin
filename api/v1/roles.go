package v1

import (
	"fmt"
	"gin-shop-admin/models"
	"gin-shop-admin/models/request"
	"gin-shop-admin/models/response"
	"gin-shop-admin/pkg/app"
	"gin-shop-admin/pkg/e"
	"gin-shop-admin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRoles(c *gin.Context) {

	roles, err := service.GetAllRoles()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_ROLES_FAIL, nil)
		return
	}

	rRoles := make([]response.RRoles, 0, len(roles))

	for _, v := range roles {
		baseRights := make([]response.Menus, 0, len(v.Rights))
		for _, right := range v.Rights {
			menus, err := service.GetRightsById(right.ID)
			if err != nil {
				app.Response(c, http.StatusInternalServerError, e.ERROR_GET_ROLES_FAIL, nil)
				return
			}
			baseRights = append(baseRights, menus)

		}
		rRole := response.RRoles{
			ID:       v.ID,
			RoleName: v.RoleName,
			RoleDesc: v.RoleDesc,
			Rights:   baseRights,
		}

		rRoles = append(rRoles, rRole)
	}

	app.Response(c, http.StatusOK, e.SUCCESS,
		rRoles,
	)
}

func AddRoles(c *gin.Context) {
	var addRoles request.AddRoles
	err := c.ShouldBindJSON(&addRoles)
	if err != nil {
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	_, exist, err := service.GetRoleByRoleName(addRoles.RoleName)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_ROLE_FAIL, nil)
		return
	}
	if exist == true {
		app.Response(c, http.StatusOK, e.ERROR_ADD_EXIST_ROLE, nil)
		return
	}

	newRole := models.Roles{
		RoleName: addRoles.RoleName,
		RoleDesc: addRoles.RoleDesc,
	}
	err = service.CreateRoles(newRole)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_ROLE_FAIL, nil)
		return
	}

	role, exist, err := service.GetRoleByRoleName(newRole.RoleName)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_ROLE_FAIL, nil)
		return
	}
	if exist == false {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_ROLE_FAIL, nil)
		return
	}

	app.Response(c, http.StatusCreated, e.CREATED_SUCCESS, gin.H{
		"roleId":   role.ID,
		"roleame":  role.RoleName,
		"roleDesc": role.RoleDesc,
	})
}

func GetRoleByID(c *gin.Context) {
	rid := c.Param("id")

	//todo ：需要验证数据正确性

	id, err := strconv.Atoi(rid)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_ROLE_FAIL, nil)
		return
	}

	role, exist, err := service.GetRoleByID(uint(id))
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_USER_FAIL, nil)
		return
	}

	if exist == false {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, gin.H{
		"roleId":   role.ID,
		"roleame":  role.RoleName,
		"roleDesc": role.RoleDesc,
	})
}

func EditRole(c *gin.Context) {
	rid := c.Param("id")
	var rRole request.AddRoles
	err := c.ShouldBind(&rRole)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPDATE_ROLE_FAIL, nil)
		return
	}

	//todo ：需要验证数据正确性

	id, err := strconv.Atoi(rid)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPDATE_ROLE_FAIL, nil)
		return
	}

	role, exist, err := service.GetRoleByID(uint(id))
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPDATE_ROLE_FAIL, nil)
		return
	}
	if exist == false {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_ROLE, nil)
		return
	}

	role.RoleName = rRole.RoleName
	role.RoleDesc = rRole.RoleDesc

	err = service.UpdateRole(role)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_UPDATE_ROLE_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, gin.H{
		"roleId":   role.ID,
		"roleame":  role.RoleName,
		"roleDesc": role.RoleDesc,
	})
}

func DeleteRole(c *gin.Context) {
	rid := c.Param("id")

	//todo ：需要验证数据正确性

	id, err := strconv.Atoi(rid)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_ROLE_FAIL, nil)
		return
	}

	_, exist, err := service.GetRoleByID(uint(id))
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_ROLE_FAIL, nil)
		return
	}
	if exist == false {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_ROLE, nil)
		return
	}

	err = service.DeleteRole(uint(id))
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_ROLE_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func AllotRights(c *gin.Context) {
	rid := c.Param("roleId")
	var rids request.AllotRids
	err := c.ShouldBindJSON(&rids)
	if err != nil {
		app.Response(c, http.StatusOK, e.ERROR_ALLOTRIGHTS_FAIL, nil)
		return
	}
	fmt.Printf(rid)
	// app.Response(rid)
}

func RemoveRight(c *gin.Context) {

}
