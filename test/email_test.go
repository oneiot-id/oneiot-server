package test

import (
	"oneiot-server/email"
	"oneiot-server/model/entity"
	"testing"
)

func TestSimpleEmailVerification(t *testing.T) {
	mail := email.Email{}

	to := entity.User{}

	//ToDo: dont use the blank identifier to discard the value but use the real one, like a real men
	_, err := mail.SendVerificationEmail(to)

	if err != nil {
		return
	}
}
