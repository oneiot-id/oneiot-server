package request

import "oneiot-server/model/entity"

type GeneralVerificationResponse struct {
	Message string                         `json:"message" :"message"`
	Payload GeneralVerificationCodePayload `json:"payload" :"payload"`
}

type GeneralVerificationCodePayload struct {
	UniqueCode     string      `json:"unique_code"`
	ExpireTimeUnix int64       `json:"expire_time_unix"`
	User           entity.User `json:"user"`
}
