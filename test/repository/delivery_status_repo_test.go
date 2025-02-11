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

func deliveryStatusRepoTestBase() repository.IDeliveryStatusRepository {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil
	}

	db := database.NewSqlConnection()
	return repository.NewDeliveryStatusRepository(db)
}

func TestCreateDeliveryStatus(t *testing.T) {
	base := deliveryStatusRepoTestBase()

	deliveryStatus, err := base.Create(context.Background(), entity.DeliveryStatuses{
		DeliveryDate:      time.Now(),
		ArrivalEstimation: time.Now(),
		RecipientName:     "John Doe",
		Courier:           "FedEx",
		Address:           "123 Main St, City, Country",
		TrackingNumber:    "TRK123456",
		DeliveryCourier:   "FedEx Express",
	})

	fmt.Println(deliveryStatus)
	assert.NoError(t, err)
	assert.NotZero(t, deliveryStatus.Id)
}

func TestGetDeliveryStatusById(t *testing.T) {
	base := deliveryStatusRepoTestBase()
	deliveryStatus, err := base.GetById(context.Background(), 1)

	fmt.Println(deliveryStatus)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), deliveryStatus.Id)
}

func TestUpdateDeliveryStatus(t *testing.T) {
	base := deliveryStatusRepoTestBase()
	deliveryStatus, err := base.Update(context.Background(), entity.DeliveryStatuses{
		Id:              1,
		RecipientName:   "Jane Doe",
		Courier:         "DHL",
		Address:         "456 Another St, City, Country",
		TrackingNumber:  "TRK654321",
		DeliveryCourier: "DHL Express",
	})

	fmt.Println(deliveryStatus)
	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", deliveryStatus.RecipientName)
}

func TestDeleteDeliveryStatus(t *testing.T) {
	base := deliveryStatusRepoTestBase()
	err := base.Delete(context.Background(), entity.DeliveryStatuses{Id: 1})

	assert.NoError(t, err)
}
