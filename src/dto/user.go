package dto

type UserSignupReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
