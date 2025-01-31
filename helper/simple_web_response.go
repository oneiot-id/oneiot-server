package helper

import "oneiot-server/response"

func ToSimpleWebResponse(message string, whatever []interface{}) response.SimpleResponse {
	return response.SimpleResponse{
		Message: message,
		Data:    whatever,
	}
}
