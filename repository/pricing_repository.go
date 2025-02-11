package repository

import (
	"context"
	"database/sql"
	"errors"
	"oneiot-server/model/entity"
	"strconv"
)

type IPricingRepository interface {
	Create(ctx context.Context, pricing entity.Pricing) (entity.Pricing, error)
	GetById(ctx context.Context, id int64) (entity.Pricing, error)
	UpdateById(ctx context.Context, pricing entity.Pricing) (entity.Pricing, error)
	DeleteById(ctx context.Context, id int64) error
}

type PricingRepository struct {
	db  *sql.DB
	ppn float64
}

func (repository *PricingRepository) Create(ctx context.Context, pricing entity.Pricing) (entity.Pricing, error) {

	//ToDo: calculate the total price
	repository.calculateTaxAndTotal(&pricing)

	query := "INSERT INTO Pricings(BasePrice, ServicePrice, DeliveryFee, Tax, AdditionalPrice, TotalPrice, PaymentsId)" +
		"VALUES(?, ?, ?, ?, ?, ?, ?) "

	execContext, err := repository.db.ExecContext(ctx, query, pricing.BasePrice, pricing.ServicePrice, pricing.DeliveryFee, pricing.Tax, pricing.AdditionalPrice, pricing.TotalPrice, pricing.PaymentsId)

	if err != nil {
		return entity.Pricing{}, errors.New("Terjadi kesalahan saat membuat pricing")
	}

	pricing.Id, _ = execContext.LastInsertId()

	return pricing, nil
}

func (repository *PricingRepository) GetById(ctx context.Context, id int64) (entity.Pricing, error) {
	var pricing = entity.Pricing{}

	query := "SELECT Id, BasePrice, ServicePrice, DeliveryFee, Tax, AdditionalPrice, TotalPrice, PaymentsId FROM Pricings WHERE Id = ?"

	err := repository.db.QueryRowContext(ctx, query, id).Scan(&pricing.Id, &pricing.BasePrice, &pricing.ServicePrice, &pricing.DeliveryFee, &pricing.Tax, &pricing.AdditionalPrice, &pricing.TotalPrice, &pricing.PaymentsId)

	if err != nil {
		return entity.Pricing{}, errors.New("Item yang dicari pada Pricings dengan id " + strconv.FormatInt(id, 10) + " tidak tersedia atau tidak ada ")
	}

	return pricing, nil
}

func (repository *PricingRepository) calculateTaxAndTotal(pricing *entity.Pricing) {
	var totalPrice = pricing.BasePrice + pricing.ServicePrice + pricing.DeliveryFee + pricing.AdditionalPrice
	var totalTax = repository.ppn * totalPrice

	pricing.Tax = totalTax
	pricing.TotalPrice = totalTax + totalPrice
}

func (repository *PricingRepository) UpdateById(ctx context.Context, pricing entity.Pricing) (entity.Pricing, error) {
	//ToDo calculate the tax and total price
	repository.calculateTaxAndTotal(&pricing)

	query := "UPDATE Pricings set BasePrice = ?, ServicePrice = ?, DeliveryFee = ?, Tax = ?, AdditionalPrice = ?, TotalPrice = ? WHERE Id = ?"

	_, err := repository.db.ExecContext(ctx, query, pricing.BasePrice, pricing.ServicePrice, pricing.DeliveryFee, pricing.Tax, pricing.AdditionalPrice, pricing.TotalPrice, pricing.Id)

	if err != nil {
		return entity.Pricing{}, errors.New("Terjadi kesalahan saat mengupdate Pricing dengan id " + strconv.FormatInt(pricing.Id, 10))
	}

	return pricing, nil
}

func (repository *PricingRepository) DeleteById(ctx context.Context, id int64) error {
	query := "DELETE FROM Pricings WHERE Id = ?"

	_, err := repository.db.ExecContext(ctx, query, id)

	if err != nil {
		return errors.New("Terjadi kesalahan saat menghapus pricing dengan id " + strconv.FormatInt(id, 10))
	}
	return nil
}

func NewPricingRepository(db *sql.DB, ppn float64) IPricingRepository {
	return &PricingRepository{db, ppn}
}
