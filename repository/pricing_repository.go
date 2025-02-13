package repository

import (
	"context"
	"database/sql"
	"errors"
	"oneiot-server/model/entity"
	"strconv"
)

type IPricingRepository interface {
	Create(ctx context.Context, tx *sql.Tx, pricing entity.Pricing) (entity.Pricing, error)
	GetById(ctx context.Context, tx *sql.Tx, id int64) (entity.Pricing, error)
	UpdateById(ctx context.Context, tx *sql.Tx, pricing entity.Pricing) (entity.Pricing, error)
	DeleteById(ctx context.Context, tx *sql.Tx, id int64) error
}

type PricingRepository struct {
	db  *sql.DB
	ppn float64
}

func (repository *PricingRepository) Create(ctx context.Context, tx *sql.Tx, pricing entity.Pricing) (entity.Pricing, error) {
	repository.calculateTaxAndTotal(&pricing)

	query := "INSERT INTO Pricings(BasePrice, ServicePrice, DeliveryFee, Tax, AdditionalPrice, TotalPrice, PaymentsId) VALUES(?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, query, pricing.BasePrice, pricing.ServicePrice, pricing.DeliveryFee, pricing.Tax, pricing.AdditionalPrice, pricing.TotalPrice, pricing.PaymentsId)
	if err != nil {
		return entity.Pricing{}, errors.New("Terjadi kesalahan saat membuat pricing")
	}

	pricing.Id, _ = result.LastInsertId()
	return pricing, nil
}

func (repository *PricingRepository) GetById(ctx context.Context, tx *sql.Tx, id int64) (entity.Pricing, error) {
	var pricing entity.Pricing
	query := "SELECT Id, BasePrice, ServicePrice, DeliveryFee, Tax, AdditionalPrice, TotalPrice, PaymentsId FROM Pricings WHERE Id = ?"

	err := tx.QueryRowContext(ctx, query, id).Scan(&pricing.Id, &pricing.BasePrice, &pricing.ServicePrice, &pricing.DeliveryFee, &pricing.Tax, &pricing.AdditionalPrice, &pricing.TotalPrice, &pricing.PaymentsId)
	if err != nil {
		return entity.Pricing{}, errors.New("Item yang dicari pada Pricings dengan id " + strconv.FormatInt(id, 10) + " tidak tersedia atau tidak ada")
	}

	return pricing, nil
}

func (repository *PricingRepository) UpdateById(ctx context.Context, tx *sql.Tx, pricing entity.Pricing) (entity.Pricing, error) {
	repository.calculateTaxAndTotal(&pricing)

	query := "UPDATE Pricings SET BasePrice = ?, ServicePrice = ?, DeliveryFee = ?, Tax = ?, AdditionalPrice = ?, TotalPrice = ? WHERE Id = ?"
	_, err := tx.ExecContext(ctx, query, pricing.BasePrice, pricing.ServicePrice, pricing.DeliveryFee, pricing.Tax, pricing.AdditionalPrice, pricing.TotalPrice, pricing.Id)

	if err != nil {
		return entity.Pricing{}, errors.New("Terjadi kesalahan saat mengupdate Pricing dengan id " + strconv.FormatInt(pricing.Id, 10))
	}

	return pricing, nil
}

func (repository *PricingRepository) DeleteById(ctx context.Context, tx *sql.Tx, id int64) error {
	query := "DELETE FROM Pricings WHERE Id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return errors.New("Terjadi kesalahan saat menghapus pricing dengan id " + strconv.FormatInt(id, 10))
	}
	return nil
}

func (repository *PricingRepository) calculateTaxAndTotal(pricing *entity.Pricing) {
	totalPrice := pricing.BasePrice + pricing.ServicePrice + pricing.DeliveryFee + pricing.AdditionalPrice
	pricing.Tax = repository.ppn * totalPrice
	pricing.TotalPrice = pricing.Tax + totalPrice
}

func NewPricingRepository(db *sql.DB, ppn float64) IPricingRepository {
	return &PricingRepository{db: db, ppn: ppn}
}
