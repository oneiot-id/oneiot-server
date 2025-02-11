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
)

func pricingRepoTestBase() repository.IPricingRepository {
	err := godotenv.Load("../../.env")

	if err != nil {
		return nil
	}

	db := database.NewSqlConnection()

	pricingRepository := repository.NewPricingRepository(db, 0.11)

	return pricingRepository
}

func TestCreatePricing(t *testing.T) {
	base := pricingRepoTestBase()

	pricing, err := base.Create(context.Background(), entity.Pricing{
		BasePrice:       1000,
		ServicePrice:    1000,
		DeliveryFee:     1000,
		AdditionalPrice: 1000,
		PaymentsId:      3,
	})

	fmt.Println(pricing)
	assert.NoError(t, err)
}

func TestUpdatePricing(t *testing.T) {
	base := pricingRepoTestBase()

	pricing, err := base.UpdateById(context.Background(), entity.Pricing{
		Id:              2,
		BasePrice:       2000,
		ServicePrice:    2000,
		DeliveryFee:     2000,
		AdditionalPrice: 2000,
	})

	fmt.Println(pricing)
	assert.NoError(t, err)
}

func TestDeletePricing(t *testing.T) {
	base := pricingRepoTestBase()
	err := base.DeleteById(context.Background(), 2)

	assert.NoError(t, err)
}

func TestGetPricing(t *testing.T) {
	base := pricingRepoTestBase()
	pricing, err := base.GetById(context.Background(), 3)

	fmt.Println(pricing)
	assert.NoError(t, err)
}
