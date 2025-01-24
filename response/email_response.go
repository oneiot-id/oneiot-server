package response

import (
	"oneiot-server/model/entity"
)

type EmailVerificationResponse struct {
	Messsage string                `json:"message"`
	Payload  EmailVerificationBody `json:"payload"`
}

type EmailVerificationBody struct {
	UniqueCode     string      `json:"uniqueCode"`
	ExpireTimeUnix int64       `json:"expireTimeUnix"`
	User           entity.User `json:"user"`
}
