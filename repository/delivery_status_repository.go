package repository

import (
	"context"
	"database/sql"
	"errors"
	"oneiot-server/helper"
	"oneiot-server/model/entity"
	"strconv"
)

type IDeliveryStatusRepository interface {
	Create(ctx context.Context, deliveryStatus entity.DeliveryStatuses) (entity.DeliveryStatuses, error)
	GetById(ctx context.Context, id int64) (entity.DeliveryStatuses, error)
	Update(ctx context.Context, deliveryStatus entity.DeliveryStatuses) (entity.DeliveryStatuses, error)
	Delete(ctx context.Context, deliveryStatus entity.DeliveryStatuses) error
}

type DeliveryStatusRepository struct {
	db *sql.DB
}

func (repository *DeliveryStatusRepository) Create(ctx context.Context, deliveryStatus entity.DeliveryStatuses) (entity.DeliveryStatuses, error) {
	query := "INSERT INTO DeliveryStatuses(DeliveryDate, ArriveEstimation, RecipientName, Courier, Address, TrackingNumber, DeliveryCourier) VALUES(?, ?, ?, ?, ?, ?, ?)"

	execContext, err := repository.db.ExecContext(ctx, query, helper.ConvertToDateString(deliveryStatus.DeliveryDate), helper.ConvertToDateString(deliveryStatus.ArrivalEstimation), deliveryStatus.RecipientName, deliveryStatus.Courier, deliveryStatus.Address, deliveryStatus.TrackingNumber, deliveryStatus.DeliveryCourier)

	if err != nil {
		return entity.DeliveryStatuses{}, err
	}

	deliveryStatus.Id, _ = execContext.LastInsertId()

	return deliveryStatus, nil
}

func (repository *DeliveryStatusRepository) GetById(ctx context.Context, id int64) (entity.DeliveryStatuses, error) {
	var deliveryStatus entity.DeliveryStatuses

	query := "SELECT Id, DeliveryDate, ArriveEstimation, RecipientName, Courier, Address, TrackingNumber, DeliveryCourier FROM DeliveryStatuses WHERE ID = ?"

	var deliveryDate string
	var deliveryArriveEstimation string

	err := repository.db.QueryRowContext(ctx, query, id).Scan(&deliveryStatus.Id, &deliveryDate, &deliveryArriveEstimation, &deliveryStatus.RecipientName, &deliveryStatus.Courier, &deliveryStatus.Address, &deliveryStatus.TrackingNumber, &deliveryStatus.DeliveryCourier)

	if err != nil {
		return entity.DeliveryStatuses{}, errors.New("Terjadi kesalahan saat memuat DeliveryStatus dengan id " + strconv.FormatInt(id, 10))
	}

	deliveryStatus.DeliveryDate = helper.StringToDate(deliveryDate)
	deliveryStatus.ArrivalEstimation = helper.StringToDate(deliveryArriveEstimation)

	return deliveryStatus, nil
}

func (repository *DeliveryStatusRepository) Update(ctx context.Context, deliveryStatus entity.DeliveryStatuses) (entity.DeliveryStatuses, error) {
	query := "UPDATE DeliveryStatuses SET DeliveryDate = ?, ArriveEstimation = ?, RecipientName = ?, Courier = ?, Address = ?, TrackingNumber = ?, DeliveryCourier = ? WHERE ID = ?"

	_, err := repository.db.ExecContext(ctx, query, helper.ConvertToDateString(deliveryStatus.DeliveryDate), helper.ConvertToDateString(deliveryStatus.ArrivalEstimation), deliveryStatus.RecipientName, deliveryStatus.Courier, deliveryStatus.Address, deliveryStatus.TrackingNumber, deliveryStatus.DeliveryCourier, deliveryStatus.Id)

	if err != nil {
		return entity.DeliveryStatuses{}, err
	}

	return deliveryStatus, nil
}

func (repository *DeliveryStatusRepository) Delete(ctx context.Context, deliveryStatus entity.DeliveryStatuses) error {
	query := "DELETE FROM DeliveryStatuses WHERE ID = ?"
	_, err := repository.db.ExecContext(ctx, query, deliveryStatus.Id)

	if err != nil {
		return errors.New("Terjadi kesalahan saat menghapus DeliveryStatus dengan id " + strconv.FormatInt(deliveryStatus.Id, 10))
	}
	return nil
}

func NewDeliveryStatusRepository(db *sql.DB) IDeliveryStatusRepository {
	return &DeliveryStatusRepository{db: db}
}
