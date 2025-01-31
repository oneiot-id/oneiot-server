package test

import (
	"github.com/joho/godotenv"
	"oneiot-server/database"
	"testing"
)

func TestGetDatabaseConnection(t *testing.T) {
	err := godotenv.Load("../.env")

	if err != nil {
		return
	}

	_ = database.NewSqlConnection()

}
