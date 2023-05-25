package dto

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
