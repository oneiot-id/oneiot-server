package request

import (
	"oneiot-server/model/entity"
)

type UserRegisterRequest struct {
	FullName    string `json:"full_name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Picture     string `json:"picture"`
	Address     string `json:"address" validate:"required"`
	Location    string `json:"location" validate:"required"`
}

type UserLoginRequest struct {
	User entity.User `json:"user"`
}
