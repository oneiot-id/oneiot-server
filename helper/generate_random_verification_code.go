package helper

import (
	"math/rand"
	"strconv"
)

func GenerateRandomVerificationCode(length int) string {
	uniqueCode := ""

	for i := 0; i < length; i++ {
		random := strconv.Itoa(rand.Intn(10))
		uniqueCode += random
	}

	return uniqueCode
}
