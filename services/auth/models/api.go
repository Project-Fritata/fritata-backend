package models

type RegisterReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRes struct {
	Message string `json:"message"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	Message string `json:"message"`
}

type LogoutRes struct {
	Message string `json:"message"`
}
