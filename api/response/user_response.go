package response

type UserResponse struct {
	User  User  `json:"user"`
	Login Login `json:"login"`
}

type User struct {
	UserId      string `json:"user_id"`
	MailAddress string `json:"mail_address"`
}

type Login struct {
	Password string `json:"password"`
}
