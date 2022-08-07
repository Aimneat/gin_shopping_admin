package request

// User login structure
type Login struct {
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

type Register struct {
	InvitationCode string `json:"invitationCode"`
	Login
}

type AddUser struct {
	Login
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

type EditUser struct {
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}
