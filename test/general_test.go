package test

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func TestBcryptEncode(t *testing.T) {
	password, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(password, []byte("password"))

	if err != nil {
		return
	}
	//$2a$10$.l2OnsPOn6u39J/WcOD2oeqCS/6R.J8lWAVSU6FWpOUdng9OnYS2i
	//$2a$10$.VxwjLP9UF3DCrsy60P1pOMYG9esdWo2fCONlrQcBDQ/zkXsh0bQO

	fmt.Println(string(password))

}

func TestBcryptDecode(t *testing.T) {
	err := bcrypt.CompareHashAndPassword([]byte(""), []byte("password"))

	fmt.Println(err)
}
