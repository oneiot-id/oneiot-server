package repository

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"oneiot-server/database"
	"oneiot-server/model/entity"
	"oneiot-server/repository"
	"testing"
	"time"
)

func transactionRepoTestBase() repository.ITransactionRepository {
	err := godotenv.Load("../../.env")

	if err != nil {
		return nil
	}

	db := database.NewSqlConnection()

	repo := repository.NewTransactionRepository(db)

	return repo
}

func TestTransactionCreate(t *testing.T) {
	repo := transactionRepoTestBase()

	transaction, err := repo.Create(context.Background(), entity.Transaction{
		UserId:               15,
		OrderId:              10,
		PricingId:            3,
		ProductionStatusesId: 1,
		DeliveryStatusesId:   2,
		Status:               "Paid",
		CreatedAt:            time.Now(),
		Complained:           false,
	})

	fmt.Println(transaction, err)
}

func TestTransactionGet(t *testing.T) {
	repo := transactionRepoTestBase()

	transaction, err := repo.GetById(context.Background(), 1)

	fmt.Println(transaction, err)
}

func TestTransactionUpdate(t *testing.T) {
	repo := transactionRepoTestBase()
	transaction, err := repo.Update(context.Background(), entity.Transaction{
		Id:                   1,
		UserId:               15,
		OrderId:              10,
		PricingId:            3,
		ProductionStatusesId: 1,
		DeliveryStatusesId:   2,
		Status:               "Paid",
		CreatedAt:            time.Now(),
		Complained:           false,
	})

	fmt.Println(transaction, err)
}

func TestTransactionDelete(t *testing.T) {
	repo := transactionRepoTestBase()
	err := repo.Delete(context.Background(), 1)
	fmt.Println(err)
}

func TestTransactionGetByUserId(t *testing.T) {
	repo := transactionRepoTestBase()
	transactions, err := repo.GetByUserId(context.Background(), 15)

	fmt.Println(transactions, err)
}
