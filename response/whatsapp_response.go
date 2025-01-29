package response

import "oneiot-server/model/entity"

type WhatsappResponse struct {
	Message string      `json:"message"`
	Payload entity.User `json:"payload"`
}
