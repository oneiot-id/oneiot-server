package request

import "oneiot-server/model/entity"

type WhatsappVerificationRequest struct {
	User entity.User `json:"user"`
}
