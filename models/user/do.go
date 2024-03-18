package user

type SignUpRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
}

type SignInRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
}

type SignInRespond struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Gender   string `json:"gender"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}
