package service

import (
	"context"
	"database/sql"
	"oneiot-server/model/dto"
	"oneiot-server/model/entity"
	"oneiot-server/repository"
)

type ITransactionService interface {
	CreateTransaction(ctx context.Context, transactionDto dto.TransactionDto) (dto.TransactionDto, error)
	DeleteTransaction(ctx context.Context, transactionId int64) error
	GetTransaction(ctx context.Context, transaction entity.Transaction) (dto.TransactionDto, error)
	GetAllUserTransactions(ctx context.Context, userId int64) ([]dto.TransactionDto, error)
	UpdateTransaction(ctx context.Context, transactionDto dto.TransactionDto) (dto.TransactionDto, error)
}

type TransactionService struct {
	db                         *sql.DB
	paymentRepository          repository.IPaymentRepository
	pricingRepository          repository.IPricingRepository
	productionStatusRepository repository.IProductionStatusRepository
	deliveryStatusRepository   repository.IDeliveryStatusRepository
	transactionRepository      repository.ITransactionRepository
}

func (service *TransactionService) CreateTransaction(ctx context.Context, transactionDto dto.TransactionDto) (dto.TransactionDto, error) {
	//First create the payments
	var err error

	tx, err := service.db.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return dto.TransactionDto{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	transactionDto.Payment, err = service.paymentRepository.Create(ctx, tx, transactionDto.Payment)

	if err != nil {
		return dto.TransactionDto{}, err
	}

	//Second create the pricing
	transactionDto.Pricing.PaymentsId = transactionDto.Payment.Id
	transactionDto.Pricing, err = service.pricingRepository.Create(ctx, tx, transactionDto.Pricing)

	if err != nil {
		return dto.TransactionDto{}, err
	}
	//Third create the production statuses
	transactionDto.ProductionStatus, err = service.productionStatusRepository.Create(ctx, tx, transactionDto.ProductionStatus)

	if err != nil {
		return dto.TransactionDto{}, err
	}

	//then create the delivery statuses
	transactionDto.DeliveryStatus, err = service.deliveryStatusRepository.Create(ctx, tx, transactionDto.DeliveryStatus)

	if err != nil {
		return dto.TransactionDto{}, err
	}

	//last create the transaction
	transactionDto.Transaction.PricingId = transactionDto.Pricing.Id
	transactionDto.Transaction.DeliveryStatusesId = transactionDto.DeliveryStatus.Id
	transactionDto.Transaction.ProductionStatusesId = transactionDto.ProductionStatus.Id

	transactionDto.Transaction, err = service.transactionRepository.Create(ctx, tx, transactionDto.Transaction)

	if err != nil {
		return dto.TransactionDto{}, err
	}

	err = tx.Commit()

	if err != nil {
		return dto.TransactionDto{}, err
	}

	return transactionDto, nil
}

func (service *TransactionService) DeleteTransaction(ctx context.Context, transactionId int64) error {
	tx, err := service.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = service.transactionRepository.Delete(ctx, tx, transactionId)

	if err != nil {
		return err
	}

	return nil
}

func (service *TransactionService) GetTransaction(ctx context.Context, transaction entity.Transaction) (dto.TransactionDto, error) {
	var transactionDto dto.TransactionDto
	var err error

	tx, err := service.db.BeginTx(ctx, &sql.TxOptions{})

	//Get the transaction first
	transaction, err = service.transactionRepository.GetById(ctx, tx, transaction.Id)
	transactionDto.Transaction = transaction

	//Get the pricing
	transactionDto.Pricing, err = service.pricingRepository.GetById(ctx, tx, transaction.PricingId)

	if err != nil {
		return dto.TransactionDto{}, err
	}

	//Get the payment by pricing id
	transactionDto.Payment, err = service.paymentRepository.GetById(ctx, tx, transactionDto.Pricing.PaymentsId)

	if err != nil {
		return dto.TransactionDto{}, err
	}

	transactionDto.DeliveryStatus, err = service.deliveryStatusRepository.GetById(ctx, tx, transaction.DeliveryStatusesId)

	if err != nil {
		return dto.TransactionDto{}, err
	}

	transactionDto.ProductionStatus, err = service.productionStatusRepository.GetById(ctx, tx, transaction.ProductionStatusesId)

	if err != nil {
		return dto.TransactionDto{}, err
	}

	return transactionDto, nil
}

func (service *TransactionService) GetAllUserTransactions(ctx context.Context, userId int64) ([]dto.TransactionDto, error) {

	tx, err := service.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

	transactions, err := service.transactionRepository.GetByUserId(ctx, tx, userId)

	if err != nil {
		return nil, err
	}

	var transactionDtos []dto.TransactionDto

	for _, transaction := range transactions {
		//Get all
		var transactionDto dto.TransactionDto

		transactionDto.Pricing, err = service.pricingRepository.GetById(ctx, tx, transaction.PricingId)

		if err != nil {
			return nil, err
		}

		transactionDto.Payment, err = service.paymentRepository.GetById(ctx, tx, transactionDto.Pricing.PaymentsId)

		if err != nil {
			return nil, err
		}

		transactionDto.DeliveryStatus, err = service.deliveryStatusRepository.GetById(ctx, tx, transaction.DeliveryStatusesId)

		if err != nil {
			return nil, err
		}

		transactionDto.ProductionStatus, err = service.productionStatusRepository.GetById(ctx, tx, transaction.ProductionStatusesId)

		if err != nil {
			return nil, err
		}

		transactionDto.Transaction = transaction

		transactionDtos = append(transactionDtos, transactionDto)
	}

	return transactionDtos, nil
}

func (service *TransactionService) UpdateTransaction(ctx context.Context, transactionDto dto.TransactionDto) (dto.TransactionDto, error) {
	var err error

	tx, err := service.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Update the payment
	transactionDto.Payment, err = service.paymentRepository.UpdateById(ctx, tx, transactionDto.Payment)
	if err != nil {
		return dto.TransactionDto{}, err
	}

	// Update the pricing
	transactionDto.Pricing, err = service.pricingRepository.UpdateById(ctx, tx, transactionDto.Pricing)
	if err != nil {
		return dto.TransactionDto{}, err
	}

	// Update the production statuses
	transactionDto.ProductionStatus, err = service.productionStatusRepository.Update(ctx, tx, transactionDto.ProductionStatus)
	if err != nil {
		return dto.TransactionDto{}, err
	}

	// Update the delivery statuses
	transactionDto.DeliveryStatus, err = service.deliveryStatusRepository.Update(ctx, tx, transactionDto.DeliveryStatus)
	if err != nil {
		return dto.TransactionDto{}, err
	}

	// Update the transaction
	transactionDto.Transaction, err = service.transactionRepository.Update(ctx, tx, transactionDto.Transaction)
	if err != nil {
		return dto.TransactionDto{}, err
	}

	err = tx.Commit()

	if err != nil {
		return dto.TransactionDto{}, err
	}

	return transactionDto, nil
}

func NewTransactionService(db *sql.DB, transactionRepository repository.ITransactionRepository, paymentRepository repository.IPaymentRepository, pricingRepository repository.IPricingRepository, productionStatusRepository repository.IProductionStatusRepository, deliveryStatusRepository repository.IDeliveryStatusRepository) ITransactionService {
	return &TransactionService{
		db:                         db,
		transactionRepository:      transactionRepository,
		pricingRepository:          pricingRepository,
		productionStatusRepository: productionStatusRepository,
		deliveryStatusRepository:   deliveryStatusRepository,
		paymentRepository:          paymentRepository,
	}
}
