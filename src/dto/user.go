package dto

type UserSignReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
