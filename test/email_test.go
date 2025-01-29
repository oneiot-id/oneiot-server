package test

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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

func TestEmailVerification(t *testing.T) {
	err := godotenv.Load()

	if err != nil {
		return
	}

	mailBox := email.Email{}

	toEmpty := entity.User{
		FullName: "",
		Email:    "",
	}

	toErlangga := entity.User{
		FullName: "Erlangga Satrya",
		Email:    "erlanggasatrya.2021@student.uny.ac.id",
	}

	_, errToEmpty := mailBox.SendVerificationEmail(toEmpty)
	_, errToErlangga := mailBox.SendVerificationEmail(toErlangga)

	assert.NotNil(t, errToEmpty)
	assert.Nil(t, errToErlangga)
}
