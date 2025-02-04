package repository

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"oneiot-server/database"
	"oneiot-server/model/entity"
	"oneiot-server/repository"
	"testing"
	"time"
)

func orderRepositoryTestBase() repository.IOrderRepository {
	err := godotenv.Load("../../.env")

	if err != nil {
		return nil
	}

	db := database.NewSqlConnection()
	repo := repository.NewOrderRepository(db)

	return repo
}

func TestCreatingOrder(t *testing.T) {
	repo := orderRepositoryTestBase()

	order, err := repo.CreateOrder(context.Background(), entity.Order{
		UserId:        2,
		BuyerId:       2,
		OrderDetailId: 15,
		IsActive:      false,
		CreatedAt:     time.Now(),
	})

	if err != nil {
		return
	}

	assert.NoError(t, err)
	fmt.Println(order)
}

func TestGetOrderById(t *testing.T) {
	repo := orderRepositoryTestBase()

	order, err := repo.GetOrderById(context.Background(), 3)

	if err != nil {
		return
	}

	assert.NoError(t, err)
	fmt.Println(order)
}

func TestSetOrderStatus(t *testing.T) {
	repo := orderRepositoryTestBase()

	order, err := repo.SetOrderStatus(context.Background(), entity.Order{
		Id:       3,
		IsActive: true,
	})

	if err != nil {
		return
	}

	assert.NoError(t, err)
	fmt.Println(order)
}
