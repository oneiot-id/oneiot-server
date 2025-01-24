package test

import (
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/assert"
	"oneiot-server/email"
	"testing"
)

func TestGenerateEmailRandomCode(t *testing.T) {
	var uniqueCode4Digit string = email.GenerateRandomVerificationCode(4)
	var uniqueCode8Digit string = email.GenerateRandomVerificationCode(8)
	var uniqueCode12Digit string = email.GenerateRandomVerificationCode(12)
	var uniqueCode16Digit string = email.GenerateRandomVerificationCode(16)

	assert.Equal(t, len(uniqueCode4Digit), 4)
	assert.Equal(t, len(uniqueCode8Digit), 8)
	assert.Equal(t, len(uniqueCode12Digit), 12)
	assert.Equal(t, len(uniqueCode16Digit), 16)
}
