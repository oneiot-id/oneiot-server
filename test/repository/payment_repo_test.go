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

func paymentRepoTestBase() repository.IPaymentRepository {
	err := godotenv.Load("../../.env")

	if err != nil {
		return nil
	}

	db := database.NewSqlConnection()

	paymentRepository := repository.NewPaymentRepository(db)

	return paymentRepository
}

func TestCreatePayment(t *testing.T) {
	base := paymentRepoTestBase()

	payment, err := base.Create(context.Background(), entity.Payment{
		PaymentProof:     "test_proof.png",
		Invoice:          "test_invoice.pdf",
		Paid:             false,
		PaymentMethodsId: 2,
	})

	fmt.Println(payment)
	assert.NoError(t, err)
}

func TestGetPayment(t *testing.T) {
	base := paymentRepoTestBase()

	payment, err := base.GetById(context.Background(), 4)
	fmt.Println(payment)
	assert.NoError(t, err)
}

func TestUpdatePayment(t *testing.T) {
	base := paymentRepoTestBase()

	payment, err := base.UpdateById(context.Background(), entity.Payment{
		Id:           4,
		PaymentProof: "test_proof_v2.png",
		Invoice:      "test_invoice_v2.pdf",
	})

	fmt.Println(payment)
	assert.NoError(t, err)
}

func TestDeletePayment(t *testing.T) {
	base := paymentRepoTestBase()

	err := base.DeleteById(context.Background(), 4)

	assert.NoError(t, err)
}
