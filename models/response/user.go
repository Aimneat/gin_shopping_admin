package response

type Login struct {
	ID       uint   `json:"id"`
	Rid      uint   `json:"rid"` //权限id
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
