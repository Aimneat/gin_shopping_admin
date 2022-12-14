package e

const (
	ERROR_NOT_EXIST_ORDER = 10001 + iota
	ERROR_CHECK_EXIST_ORDER_FAIL
	ERROR_ADD_ORDER_FAIL
	ERROR_UPDATA_ORDER_FAIL
	ERROR_GET_ORDERS_FAIL
	ERROR_GET_ORDERS_NUMBER_FAIL
	ERROR_GET_ORDER_FAIL
	ERROR_NOT_EXIST_PRODUCT_IN_ORDER

	ERROR_NOT_EXIST_USER
	ERROR_ADD_EXIST_USER
	ERROR_ADD_USER_FAIL
	ERROR_DELETE_USER_FAIL
	ERROR_GET_USER_FAIL
	ERROR_GET_USERS_FAIL
	ERROR_GET_USER_TELEPHONE_FAIL
	ERROR_UPDATE_USER_FAIL
	ERROR_USER_INVITATIONCODE

	ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	ERROR_GENERATE_TOKEN_FAIL

	ERROR_GET_MENUS_FAIL
	ERROR_GET_RIGHTSLIST_FAIL

	ERROR_GET_ROLES_FAIL
	ERROR_GET_ROLE_FAIL
	ERROR_ADD_ROLE_FAIL
	ERROR_UPDATE_ROLE_FAIL
	ERROR_ADD_EXIST_ROLE
	ERROR_NOT_EXIST_ROLE
	ERROR_ALLOTRIGHTS_FAIL
	ERROR_DELETE_ROLE_FAIL

	ERROR_GET_CATEGORIES_FAIL

	SUCCESS         = 200
	CREATED_SUCCESS = 201
	ERROR           = 500
	INVALID_PARAMS  = 400
)
