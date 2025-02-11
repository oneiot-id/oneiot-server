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

func productionStatusRepoTestBase() repository.IProductionStatusRepository {
	err := godotenv.Load("../../.env")

	if err != nil {
		return nil
	}

	db := database.NewSqlConnection()

	repo := repository.NewProductionStatusRepository(db)

	return repo
}

func TestCreateNewProductionStatus(t *testing.T) {
	repo := productionStatusRepoTestBase()

	productionStatus, err := repo.Create(context.Background(), entity.ProductionStatus{
		ProductionDate:   time.Now(),
		EstimatedDate:    time.Now(),
		LatestStatus:     "",
		ProductionStages: "",
	})

	fmt.Println(productionStatus)
	assert.Nil(t, err)
}

func TestGetProductionStatus(t *testing.T) {
	repo := productionStatusRepoTestBase()
	productionStatus, err := repo.GetById(context.Background(), 1)

	fmt.Println(productionStatus, err)
}

func TestUpdateProductionStatus(t *testing.T) {
	repo := productionStatusRepoTestBase()

	updated, err := repo.Update(context.Background(), entity.ProductionStatus{
		Id:               1,
		ProductionDate:   time.Now(),
		EstimatedDate:    time.Now(),
		ProductionStages: "Haha",
		LatestStatus:     "HihiJson",
	})

	fmt.Println(updated, err)
}

func TestDeleteProductionStatus(t *testing.T) {
	repo := productionStatusRepoTestBase()

	err := repo.DeleteById(context.Background(), 2)

	fmt.Println(err)
}
