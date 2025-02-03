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

func orderDetailsRepositoryTestBase() repository.IOrderDetailRepository {
	err := godotenv.Load("../../.env")

	if err != nil {
		return nil
	}

	db := database.NewSqlConnection()
	orderDetailsRepo := repository.NewOrderDetailRepository(db)

	return orderDetailsRepo
}

func TestCreatingOrderDetails(t *testing.T) {
	repo := orderDetailsRepositoryTestBase()

	regularText := entity.OrderSpeed(entity.Regular).String()
	//timeNow := time.Now().Format("06.01.02")

	detail, err := repo.CreateOrderDetail(context.Background(),
		entity.OrderDetail{
			OrderName:        "Peralatan Tangan",
			ServicesId:       1,
			Deadline:         time.Now(),
			Speed:            regularText,
			BriefFile:        "brieffile.pdf",
			ImportantPoint:   "jsonfile",
			AdditionalNotes:  "additional notes",
			OrderSummaryFile: "summaryfile.pdf",
		})

	if err != nil {
		return
	}

	assert.Nil(t, err)

	fmt.Println(detail)
}

func TestDeleteOrderDetails(t *testing.T) {
	repo := orderDetailsRepositoryTestBase()

	err := repo.DeleteOrderDetail(context.Background(), entity.OrderDetail{
		Id: 16,
	})
	if err != nil {
		return
	}

	assert.Nil(t, err)
}

func TestGetOrderDetailsById(t *testing.T) {
	repo := orderDetailsRepositoryTestBase()

	orderDetails, err := repo.GetOrderById(context.Background(), entity.OrderDetail{
		Id: 15,
	})
	if err != nil {
		return
	}

	fmt.Println(orderDetails)
	assert.Nil(t, err)
}
