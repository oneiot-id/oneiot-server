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

func paymentMethodRepositoryTestBase() repository.IPaymentMethodRepository {
	err := godotenv.Load("../../.env")

	if err != nil {
		return nil
	}

	db := database.NewSqlConnection()

	r := repository.NewPaymentMethodRepository(db)
	return r
}

func TestCreateNewPaymentMethod(t *testing.T) {
	base := paymentMethodRepositoryTestBase()

	paymentMethod, err := base.Create(context.Background(), entity.PaymentMethod{
		Name:    "BCA Virtual Account",
		Number:  "2501142524",
		Logo:    "localhost/bca",
		Acronym: "BCA VA",
	})

	if err != nil {
		return
	}

	fmt.Println(paymentMethod)
}

func TestDeletePaymentMethod(t *testing.T) {
	base := paymentMethodRepositoryTestBase()

	err := base.DeleteById(context.Background(), 3)

	if err != nil {
		return
	}
}

func TestGetPaymentMethodById(t *testing.T) {
	base := paymentMethodRepositoryTestBase()

	paymentMethod, err := base.GetById(context.Background(), 2)

	if err != nil {
		return
	}

	fmt.Println(paymentMethod)
}

func TestUpdatePaymentMethod(t *testing.T) {
	base := paymentMethodRepositoryTestBase()

	paymentMethod, err := base.UpdateById(context.Background(), entity.PaymentMethod{
		Id:      2,
		Name:    "QRIS OneIoT Id",
		Number:  "0823",
		Logo:    "logo",
		Acronym: "BCA VA",
	})

	fmt.Println(paymentMethod)

	assert.NoError(t, err)
}

func TestGetAllPaymentMethods(t *testing.T) {
	base := paymentMethodRepositoryTestBase()

	paymentMethods, err := base.GetAllPaymentMethods(context.Background())

	fmt.Println(paymentMethods)
	assert.NoError(t, err)
}
