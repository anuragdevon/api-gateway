package dto

import "api-gateway/pkg/auth/pb"

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var UserTypeMap = map[string]pb.UserType{
	"admin":    pb.UserType_ADMIN,
	"customer": pb.UserType_CUSTOMER,
}
