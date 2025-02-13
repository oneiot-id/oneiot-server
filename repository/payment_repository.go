package repository

import (
	"context"
	"database/sql"
	"errors"
	"oneiot-server/model/entity"
	"strconv"
)

type IPaymentRepository interface {
	Create(ctx context.Context, tx *sql.Tx, payment entity.Payment) (entity.Payment, error)
	GetById(ctx context.Context, tx *sql.Tx, id int64) (entity.Payment, error)
	DeleteById(ctx context.Context, tx *sql.Tx, id int64) error
	UpdateById(ctx context.Context, tx *sql.Tx, payment entity.Payment) (entity.Payment, error)
}

type PaymentRepository struct {
	db *sql.DB
}

func (repository *PaymentRepository) Create(ctx context.Context, tx *sql.Tx, payment entity.Payment) (entity.Payment, error) {
	query := "INSERT INTO Payments(PaymentProof, Invoice, Paid, PaymentMethodsId) VALUES(?, ?, ?, ?)"

	var execContext sql.Result
	var err error

	if tx != nil {
		execContext, err = tx.ExecContext(ctx, query, payment.PaymentProof, payment.Invoice, payment.Paid, payment.PaymentMethodsId)
	} else {
		execContext, err = repository.db.ExecContext(ctx, query, payment.PaymentProof, payment.Invoice, payment.Paid, payment.PaymentMethodsId)
	}

	if err != nil {
		return entity.Payment{}, errors.New("Terjadi kesalahan saat membuat payment")
	}

	payment.Id, _ = execContext.LastInsertId()

	return payment, nil
}

func (repository *PaymentRepository) GetById(ctx context.Context, tx *sql.Tx, id int64) (entity.Payment, error) {
	var payment entity.Payment
	query := "SELECT Id, PaymentProof, Invoice, Paid, PaymentMethodsId FROM Payments WHERE Id = ?"

	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, id).Scan(&payment.Id, &payment.PaymentProof, &payment.Invoice, &payment.Paid, &payment.PaymentMethodsId)
	} else {
		err = repository.db.QueryRowContext(ctx, query, id).Scan(&payment.Id, &payment.PaymentProof, &payment.Invoice, &payment.Paid, &payment.PaymentMethodsId)
	}

	if err != nil {
		return entity.Payment{}, errors.New("Terjadi kesalahan saat mendapatkan payment dengan id " + strconv.FormatInt(id, 10))
	}

	return payment, nil
}

func (repository *PaymentRepository) DeleteById(ctx context.Context, tx *sql.Tx, id int64) error {
	query := "DELETE FROM Payments WHERE id = ?"

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, id)
	} else {
		_, err = repository.db.ExecContext(ctx, query, id)
	}

	if err != nil {
		return errors.New("Terjadi kesalahan saat menghapus payment dengan id " + strconv.FormatInt(id, 10))
	}

	return nil
}

func (repository *PaymentRepository) UpdateById(ctx context.Context, tx *sql.Tx, payment entity.Payment) (entity.Payment, error) {
	query := "UPDATE Payments SET PaymentProof = ?, Invoice = ?, Paid = ? WHERE id = ?"

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, payment.PaymentProof, payment.Invoice, payment.Paid, payment.Id)
	} else {
		_, err = repository.db.ExecContext(ctx, query, payment.PaymentProof, payment.Invoice, payment.Paid, payment.Id)
	}

	if err != nil {
		return entity.Payment{}, errors.New("Terjadi kesalahan saat memperbarui payment dengan id " + strconv.FormatInt(payment.Id, 10))
	}

	return payment, nil
}

func NewPaymentRepository(db *sql.DB) IPaymentRepository {
	return &PaymentRepository{db: db}
}
