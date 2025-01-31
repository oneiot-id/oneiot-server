package response

import "oneiot-server/model/entity"

type UserRegisterResponse struct {
	Message string      `json:"message"`
	Data    entity.User `json:"data"`
}
