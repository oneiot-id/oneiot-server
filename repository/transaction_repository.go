package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"oneiot-server/helper"
	"oneiot-server/model/entity"
)

type ITransactionRepository interface {
	Create(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error)
	GetById(ctx context.Context, transactionId int64) (entity.Transaction, error)
	GetByUserId(ctx context.Context, userId int64) ([]entity.Transaction, error)
	Update(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error)
	Delete(ctx context.Context, transactionId int64) error
}

type TransactionRepository struct {
	db *sql.DB
}

func (repo *TransactionRepository) Create(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	query := "INSERT INTO Transactions(UserId, OrderId, PricingId, ProductionStatusesId, DeliveryStatusesId, Status, CreatedAt, Complained) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	execContext, err := repo.db.ExecContext(ctx, query, transaction.UserId, transaction.OrderId, transaction.PricingId, transaction.ProductionStatusesId, transaction.DeliveryStatusesId, transaction.Status, helper.ConvertToDateTimeString(transaction.CreatedAt), transaction.Complained)

	if err != nil {
		return entity.Transaction{}, errors.New("Terjadi kesalahan saat membuat transaction")
	}

	transaction.Id, _ = execContext.LastInsertId()

	return transaction, nil
}

func (repo *TransactionRepository) GetById(ctx context.Context, transactionId int64) (entity.Transaction, error) {
	var transaction entity.Transaction

	query := "SELECT Id, UserId, OrderId, PricingId, ProductionStatusesId, DeliveryStatusesId, Status, CreatedAt, Complained FROM Transactions WHERE Id = ?"

	var createdAt string

	err := repo.db.QueryRowContext(ctx, query, transactionId).Scan(&transaction.Id, &transaction.UserId, &transaction.ProductionStatusesId, &transaction.OrderId, &transaction.ProductionStatusesId, &transaction.DeliveryStatusesId, &transaction.Status, &createdAt, &transaction.Complained)

	if err != nil {
		return entity.Transaction{}, errors.New("Terjadi kesalahan saat mendapatkan transaction dengan id " + fmt.Sprint(transactionId))
	}

	transaction.CreatedAt = helper.StringToDateTime(createdAt)

	return transaction, nil
}

func (repo *TransactionRepository) GetByUserId(ctx context.Context, userId int64) ([]entity.Transaction, error) {
	query := "SELECT Id, UserId, OrderId, PricingId, ProductionStatusesId, DeliveryStatusesId, Status, CreatedAt, Complained FROM Transactions WHERE UserId = ?"

	rows, err := repo.db.QueryContext(ctx, query, userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []entity.Transaction

	for rows.Next() {
		transaction := entity.Transaction{}

		var createdAt string

		err := rows.Scan(&transaction.Id, &transaction.UserId, &transaction.OrderId, &transaction.PricingId, &transaction.ProductionStatusesId, &transaction.DeliveryStatusesId, &transaction.Status, &createdAt, &transaction.Complained)

		transaction.CreatedAt = helper.StringToDateTime(createdAt)

		if err != nil {
			return nil, errors.New("Terjadi kesalahan saat mendapatkan transaction dengan id " + fmt.Sprint(transaction.Id))
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (repo *TransactionRepository) Update(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	query := "UPDATE Transactions SET UserId=?, OrderId=?, PricingId=?, ProductionStatusesId=?, DeliveryStatusesId=?, Status=?, CreatedAt=?, Complained=? WHERE Id=?"

	_, err := repo.db.ExecContext(ctx, query, transaction.UserId, transaction.OrderId, transaction.PricingId,
		transaction.ProductionStatusesId, transaction.DeliveryStatusesId, transaction.Status,
		helper.ConvertToDateTimeString(transaction.CreatedAt), transaction.Complained, transaction.Id)

	if err != nil {
		return entity.Transaction{}, errors.New("Terjadi kesalahan saat mengupdate transaction dengan id " + fmt.Sprint(transaction.Id))
	}

	return transaction, nil
}

func (repo *TransactionRepository) Delete(ctx context.Context, transactionId int64) error {
	query := "DELETE FROM Transactions WHERE Id = ?"
	_, err := repo.db.ExecContext(ctx, query, transactionId)

	if err != nil {
		return errors.New("Failed to delete transaction")
	}
	return nil
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db}
}
