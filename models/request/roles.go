package request

type AddRoles struct {
	RoleName string `json:"roleName"`
	RoleDesc string `json:"roleDesc"`
}

type AllotRids struct {
	Rids string `json:"rids"`
}
