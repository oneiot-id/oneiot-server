package test

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestGenerateRandomUuid(t *testing.T) {
	newUuid, err := uuid.NewUUID()

	fmt.Println(newUuid, err)
}

func TestGenerateTime(t *testing.T) {
	timeString := time.Now().Format("2006-01-02 15-04-05")

	fmt.Println(timeString)
}
