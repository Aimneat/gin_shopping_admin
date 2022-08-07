package e

var MsgFlags = map[int]string{
	SUCCESS:         "ok",
	CREATED_SUCCESS: "创建成功",
	ERROR:           "fail",
	INVALID_PARAMS:  "请求参数错误",

	ERROR_NOT_EXIST_ORDER:            "该订单不存在",
	ERROR_CHECK_EXIST_ORDER_FAIL:     "检查订单是否存在失败",
	ERROR_ADD_ORDER_FAIL:             "新增订单失败",
	ERROR_UPDATA_ORDER_FAIL:          "更新订单失败",
	ERROR_GET_ORDERS_FAIL:            "获取多个订单失败",
	ERROR_GET_ORDERS_NUMBER_FAIL:     "获取订单数目失败",
	ERROR_GET_ORDER_FAIL:             "获取单个订单失败",
	ERROR_NOT_EXIST_PRODUCT_IN_ORDER: "订单中商品不存在",

	ERROR_NOT_EXIST_USER:          "用户不存在或密码错误",
	ERROR_ADD_EXIST_USER:          "添加已存在用户",
	ERROR_ADD_USER_FAIL:           "添加用户失败",
	ERROR_DELETE_USER_FAIL:        "删除用户失败",
	ERROR_GET_USER_FAIL:           "获取用户失败",
	ERROR_GET_USERS_FAIL:          "获取用户列表失败",
	ERROR_GET_USER_TELEPHONE_FAIL: "获取用户电话号码失败",
	ERROR_UPDATE_USER_FAIL:        "更新用户失败",
	ERROR_USER_INVITATIONCODE:     "邀请码错误",

	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token已超时",
	ERROR_GENERATE_TOKEN_FAIL:      "token生成失败",

	ERROR_GET_MENUS_FAIL:      "获取左侧菜单失败",
	ERROR_GET_RIGHTSLIST_FAIL: "获取权限列表失败",
	ERROR_GET_ROLES_FAIL:      "获取角色列表失败",
	ERROR_GET_ROLE_FAIL:       "获取角色失败",
	ERROR_ADD_ROLE_FAIL:       "添加角色错误",
	ERROR_UPDATE_ROLE_FAIL:    "修改角色错误",
	ERROR_ADD_EXIST_ROLE:      "添加已存在角色",
	ERROR_NOT_EXIST_ROLE:      "不存在的角色",
	ERROR_ALLOTRIGHTS_FAIL:    "分配权限失败",
	ERROR_DELETE_ROLE_FAIL:    "删除角色失败",

	ERROR_GET_CATEGORIES_FAIL: "获取分类列表失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
