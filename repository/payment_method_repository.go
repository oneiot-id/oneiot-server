package repository

import (
	"context"
	"database/sql"
	"errors"
	"oneiot-server/model/entity"
	"strconv"
)

type IPaymentMethodRepository interface {
	Create(ctx context.Context, paymentMethod entity.PaymentMethod) (entity.PaymentMethod, error)
	GetById(ctx context.Context, id int64) (entity.PaymentMethod, error)
	UpdateById(ctx context.Context, paymentMethod entity.PaymentMethod) (entity.PaymentMethod, error)
	DeleteById(ctx context.Context, id int64) error
	GetAllPaymentMethods(ctx context.Context) ([]entity.PaymentMethod, error)
}

type PaymentMethodRepository struct {
	db *sql.DB
}

func (repository *PaymentMethodRepository) GetAllPaymentMethods(ctx context.Context) ([]entity.PaymentMethod, error) {
	query := "SELECT Id, Name, Number, Logo, Acronym from PaymentMethods"

	rows, err := repository.db.QueryContext(ctx, query)

	if err != nil {
		return nil, errors.New("Terjadi kesalahan pada query")
	}

	defer rows.Close()

	var paymentMethods []entity.PaymentMethod

	for rows.Next() {
		var paymentMethod entity.PaymentMethod

		err := rows.Scan(&paymentMethod.Id, &paymentMethod.Name, &paymentMethod.Number, &paymentMethod.Logo, &paymentMethod.Acronym)

		if err != nil {
			return nil, errors.New("Terjadi kesalahan pada scanning payment method")
		}

		paymentMethods = append(paymentMethods, paymentMethod)
	}

	if len(paymentMethods) == 0 {
		return []entity.PaymentMethod{}, errors.New("Belum terdapat payment methods pada database")
	}

	return paymentMethods, nil
}

func (repository *PaymentMethodRepository) Create(ctx context.Context, paymentMethod entity.PaymentMethod) (entity.PaymentMethod, error) {
	query := "INSERT INTO PaymentMethods(Name, Number, Logo, Acronym)" +
		"VALUES (?, ?, ?, ?)"

	exec, err := repository.db.ExecContext(ctx, query, paymentMethod.Name, paymentMethod.Number, paymentMethod.Logo, paymentMethod.Acronym)

	if err != nil {
		return entity.PaymentMethod{}, errors.New("Terjadi kesalahan saat membuat payment method")
	}

	paymentMethod.Id, _ = exec.LastInsertId()

	return paymentMethod, nil
}

func (repository *PaymentMethodRepository) GetById(ctx context.Context, id int64) (entity.PaymentMethod, error) {
	var paymentMethod entity.PaymentMethod

	query := "SELECT Id, Name, Number, Logo, Acronym FROM PaymentMethods WHERE Id=?"

	err := repository.db.QueryRowContext(ctx, query, id).Scan(&paymentMethod.Id, &paymentMethod.Name, &paymentMethod.Number, &paymentMethod.Logo, &paymentMethod.Acronym)

	if err != nil {
		return entity.PaymentMethod{}, errors.New("Terjadi kesalahan saat mendapatkan payment method dengan id " + strconv.Itoa(int(id)))
	}

	return paymentMethod, nil
}

func (repository *PaymentMethodRepository) UpdateById(ctx context.Context, paymentMethod entity.PaymentMethod) (entity.PaymentMethod, error) {
	query := "UPDATE PaymentMethods SET Name=?, Number=?, Logo=?, Acronym=? WHERE Id=?"

	execContext, err := repository.db.ExecContext(ctx, query, paymentMethod.Name, paymentMethod.Number, paymentMethod.Logo, paymentMethod.Acronym, paymentMethod.Id)

	if err != nil {
		return entity.PaymentMethod{}, err
	}

	paymentMethod.Id, _ = execContext.LastInsertId()

	return paymentMethod, nil
}

func (repository *PaymentMethodRepository) DeleteById(ctx context.Context, id int64) error {
	query := "DELETE FROM PaymentMethods WHERE Id=?"

	_, err := repository.db.ExecContext(ctx, query, id)

	if err != nil {
		return errors.New("Terjadi kesalahan saat menghapus payment method dengan id " + strconv.FormatInt(id, 10))
	}

	return nil
}

func NewPaymentMethodRepository(db *sql.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{db: db}
}
