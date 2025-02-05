package helper

import (
	"errors"
	"oneiot-server/model/entity"
)

func ValidateUserRegister(user entity.User) error {

	//When the user_pictures field is empty
	if user.FullName == "" && user.Password == "" && user.Address == "" && user.PhoneNumber == "" && user.Email == "" {
		return errors.New("Terdapat beberapa data yang belum lengkap")
	}
	return nil
}
