package repository_test

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"oneiot-server/database"
	"oneiot-server/model/entity"
	"oneiot-server/repository"
	"testing"
)

func buyerRepositoryBase() repository.IBuyerRepository {
	err := godotenv.Load("../../.env")

	if err != nil {
		return nil
	}

	db := database.NewSqlConnection()
	repo := repository.NewBuyerRepository(db)

	return repo
}

func TestCreateBuyer(t *testing.T) {
	repo := buyerRepositoryBase()

	buyerDetails, err := repo.Create(context.Background(), entity.Buyer{
		FullName:        "Vincent Kenutama",
		Email:           "vincent@gmail.com",
		PhoneNumber:     "072112123",
		FullAddress:     "Jonggol Barat, Jakarta",
		AdditionalNotes: "Yang penting nyoba dulu hehe",
	})

	if err != nil {
		return
	}

	fmt.Println(buyerDetails)
	assert.NoError(t, err)
}

func TestGetBuyerById(t *testing.T) {
	repo := buyerRepositoryBase()

	buyerDetails, err := repo.GetById(context.Background(), entity.Buyer{
		Id: 2,
	})

	if err != nil {
		return
	}

	fmt.Println(buyerDetails)
	assert.NoError(t, err)
}
