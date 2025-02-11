package repository

import (
	"context"
	"database/sql"
	"errors"
	"oneiot-server/helper"
	"oneiot-server/model/entity"
	"strconv"
)

type IProductionStatusRepository interface {
	Create(ctx context.Context, productionStatus entity.ProductionStatus) (entity.ProductionStatus, error)
	GetById(ctx context.Context, id int64) (entity.ProductionStatus, error)
	Update(ctx context.Context, productionStatus entity.ProductionStatus) (entity.ProductionStatus, error)
	DeleteById(ctx context.Context, id int64) error
}

type ProductionStatusRepository struct {
	db *sql.DB
}

func (repository *ProductionStatusRepository) Create(ctx context.Context, productionStatus entity.ProductionStatus) (entity.ProductionStatus, error) {
	query := "INSERT INTO ProductionStatuses(ProductionDate, EstimatedDate, LatestStatus, ProductionStages)" +
		"VALUES(?, ?, ?, ?)"

	execContext, err := repository.db.ExecContext(ctx, query, helper.ConvertToDateTimeString(productionStatus.ProductionDate), helper.ConvertToDateTimeString(productionStatus.EstimatedDate), productionStatus.LatestStatus, productionStatus.ProductionStages)

	if err != nil {
		return entity.ProductionStatus{}, errors.New("Terjadi kesalahan saat membuat production status")
	}

	productionStatus.Id, _ = execContext.LastInsertId()

	return productionStatus, nil
}

func (repository *ProductionStatusRepository) GetById(ctx context.Context, id int64) (entity.ProductionStatus, error) {
	var productionStatus entity.ProductionStatus

	query := "SELECT ID, ProductionDate, EstimatedDate, LatestStatus, ProductionStages FROM ProductionStatuses WHERE ID = ?"

	var productionDate string
	var productionEstimated string

	err := repository.db.QueryRowContext(ctx, query, id).Scan(&productionStatus.Id, &productionDate, &productionEstimated, &productionStatus.LatestStatus, &productionStatus.ProductionStages)

	if err != nil {
		return entity.ProductionStatus{}, errors.New("Kesalahan saat memuat production status dengan id " + strconv.FormatInt(id, 10))
	}

	productionStatus.ProductionDate = helper.StringToDate(productionDate)
	productionStatus.EstimatedDate = helper.StringToDate(productionEstimated)

	return productionStatus, nil
}

func (repository *ProductionStatusRepository) Update(ctx context.Context, productionStatus entity.ProductionStatus) (entity.ProductionStatus, error) {
	query := "UPDATE ProductionStatuses Set ProductionDate = ?, EstimatedDate = ?, LatestStatus = ?, ProductionStages = ? WHERE Id = ?"

	_, err := repository.db.ExecContext(ctx, query, helper.ConvertToDateTimeString(productionStatus.ProductionDate), helper.ConvertToDateTimeString(productionStatus.EstimatedDate), productionStatus.LatestStatus, productionStatus.ProductionStages, productionStatus.Id)

	if err != nil {
		return entity.ProductionStatus{}, errors.New("Kesalahan saat memperbarui production status dengan id " + strconv.FormatInt(productionStatus.Id, 10))
	}

	return productionStatus, nil
}

func (repository *ProductionStatusRepository) DeleteById(ctx context.Context, id int64) error {
	query := "DELETE FROM ProductionStatuses WHERE ID = ?"
	_, err := repository.db.ExecContext(ctx, query, id)

	if err != nil {
		return errors.New("Kesalahan saat menghapus production status dengan id " + strconv.FormatInt(id, 10))
	}

	return nil
}

func NewProductionStatusRepository(db *sql.DB) IProductionStatusRepository {
	return &ProductionStatusRepository{db: db}
}
