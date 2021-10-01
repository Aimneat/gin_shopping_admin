package request

// User login structure
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	InvitationCode string `json:"invitationCode"`
	Login
}
