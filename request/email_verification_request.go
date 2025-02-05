package request

import "oneiot-server/model/entity"

type EmailVerificationRequestBody struct {
	User entity.User `json:"user_pictures"`
}
