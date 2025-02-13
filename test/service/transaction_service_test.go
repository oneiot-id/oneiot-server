package service

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"oneiot-server/database"
	"oneiot-server/helper"
	"oneiot-server/model/dto"
	"oneiot-server/model/entity"
	"oneiot-server/repository"
	"oneiot-server/service"
	"testing"
	"time"
)

func transactionServiceTestBase() service.ITransactionService {
	err := godotenv.Load("../../.env")

	if err != nil {
		return nil
	}

	db := database.NewSqlConnection()

	transactionRepo := repository.NewTransactionRepository(db)
	productionStatusRepository := repository.NewProductionStatusRepository(db)
	deliveryStatusRepository := repository.NewDeliveryStatusRepository(db)
	pricingRepository := repository.NewPricingRepository(db, 0.11)
	paymentRepository := repository.NewPaymentRepository(db)

	transactionService := service.NewTransactionService(db, transactionRepo, paymentRepository, pricingRepository, productionStatusRepository, deliveryStatusRepository)

	return transactionService
}

func TestCreateTransaction(t *testing.T) {
	transactionService := transactionServiceTestBase()

	var transaction, err = transactionService.CreateTransaction(context.Background(), dto.TransactionDto{
		Transaction: entity.Transaction{
			UserId:     15,
			OrderId:    10,
			Status:     "Unpaid",
			CreatedAt:  time.Now(),
			Complained: false,
		},
		Pricing: entity.Pricing{
			BasePrice:       2000,
			ServicePrice:    2000,
			DeliveryFee:     2000,
			AdditionalPrice: 4000,
		},
		Payment: entity.Payment{
			PaymentProof:     "proof.jpg",
			Invoice:          "invoice.pdf",
			Paid:             false,
			PaymentMethodsId: 2,
		},
		ProductionStatus: entity.ProductionStatus{
			ProductionDate:   time.Now(),
			EstimatedDate:    time.Now(),
			LatestStatus:     "Unpaid",
			ProductionStages: "{}",
		},
		DeliveryStatus: entity.DeliveryStatuses{
			DeliveryDate:      time.Now(),
			ArrivalEstimation: time.Now(),
			RecipientName:     "Erlangga",
			Courier:           "JNT",
			Address:           "Jakarta",
			TrackingNumber:    "1231",
			DeliveryCourier:   "JNT",
		},
	})

	fmt.Println(transaction, err)
}

func TestGetTransaction(t *testing.T) {
	transactionService := transactionServiceTestBase()

	transaction, err := transactionService.GetTransaction(context.Background(), entity.Transaction{
		Id: 11,
	})

	json := helper.MarshalThis(transaction)
	fmt.Println(transaction, err, json)

}

//
//func TestUpdateTransaction(t *testing.T) {
//	transactionService := transactionServiceTestBase()
//
//	transaction, err := transactionService.UpdateTransaction(context.Background(), dto.TransactionDto{
//		Transaction: entity.Transaction{
//			Id: 10
//		}
//	})
//}

func TestGetAllTransaction(t *testing.T) {
	transactionService := transactionServiceTestBase()

	transactions, err := transactionService.GetAllUserTransactions(context.Background(), 15)

	fmt.Println(transactions, err)
}

func TestDeleteTransaction(t *testing.T) {
	transactionService := transactionServiceTestBase()

	err := transactionService.DeleteTransaction(context.Background(), 10)

	fmt.Println(err)
}
