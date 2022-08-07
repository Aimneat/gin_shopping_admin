package response

type Login struct {
	ID       uint   `json:"id"`
	RolesID  uint   `json:"rid"` //权限id
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type Users struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Mobile    string `json:"mobile"`
	Type      uint   `json:"type"`
	Email     string `json:"email"`
	CreatedAt string `json:"create_time"`
	MgState   bool   `json:"mg_state"`
	RoleName  string `json:"role_name"`
}

type AddUser struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Mobile    string `json:"mobile"`
	Type      uint   `json:"type"`
	Openid    string `json:"openid"`
	Email     string `json:"email"`
	CreatedAt string `json:"create_time"`
	UpdatedAt string `json:"modify_time"`
	IsDelete  bool   `json:"is_delete"`
	IsActive  string `json:"is_active"`
}

type UserStateChanged struct {
	ID       uint   `json:"id"`
	RolesID  uint   `json:"rid"` //权限id
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	MgState  bool   `json:"mg_state"`
}
