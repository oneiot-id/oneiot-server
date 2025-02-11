package repository

import (
	"context"
	"database/sql"
	"oneiot-server/model/entity"
)

type ITransactionRepository interface {
	Create(ctx context.Context, transaction entity.Transaction, payment entity.Payment, pricing entity.Pricing)
}

type TransactionRepository struct {
	db *sql.DB
}
