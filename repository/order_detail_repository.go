package repository

import (
	"context"
	"database/sql"
	"errors"
	"oneiot-server/model/entity"
	"time"
)

type IOrderDetailRepository interface {
	CreateOrderDetail(ctx context.Context, orderDetail entity.OrderDetail) (entity.OrderDetail, error)
	DeleteOrderDetail(ctx context.Context, orderDetail entity.OrderDetail) error
	GetOrderById(ctx context.Context, orderDetail entity.OrderDetail) (entity.OrderDetail, error)

	//ToDo: Sepertinya kita harus menambahkan setter untuk order ups bukan
	//for now we cannot update the order, because the order that is posted cannot be edited
}

type OrderDetailRepository struct {
	db        *sql.DB
	tableName string
}

func (repository *OrderDetailRepository) GetOrderById(ctx context.Context, orderDetail entity.OrderDetail) (entity.OrderDetail, error) {
	query := "SELECT * FROM OrderDetails WHERE Id = ?"

	rows, err := repository.db.QueryContext(ctx, query, orderDetail.Id)

	if err != nil {
		return entity.OrderDetail{}, errors.New("Error saat mendapatkan id")
	}

	defer rows.Close()

	if !rows.Next() {
		return entity.OrderDetail{}, errors.New("Error tidak ada item order details dengan id " + string(rune(orderDetail.Id)))
	}

	var deadlineString string

	err = rows.Scan(&orderDetail.Id, &orderDetail.OrderName, &orderDetail.ServicesId, &deadlineString, &orderDetail.Speed, &orderDetail.BriefFile, &orderDetail.ImportantPoint, &orderDetail.AdditionalNotes, &orderDetail.OrderSummaryFile)

	orderDetail.Deadline, err = time.Parse("2006-01-02", deadlineString)

	if err != nil {
		return entity.OrderDetail{}, errors.New("Error saat scanning order details dengan id " + string(rune(orderDetail.Id)))
	}

	return orderDetail, nil
}

func (repository *OrderDetailRepository) DeleteOrderDetail(ctx context.Context, orderDetail entity.OrderDetail) error {
	query := "DELETE FROM OrderDetails WHERE id = ?"

	_, err := repository.db.ExecContext(ctx, query, orderDetail.Id)

	if err != nil {
		return errors.New("Error saat menghapus order detail")
	}

	return nil
}

func (repository *OrderDetailRepository) CreateOrderDetail(ctx context.Context, orderDetail entity.OrderDetail) (entity.OrderDetail, error) {
	query := "INSERT INTO OrderDetails(OrderName, ServicesId, Deadline, Speed, BriefFile, ImportantPoint, AdditionalNotes, OrderSummaryFile) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?)"

	execContext, err := repository.db.ExecContext(ctx, query, orderDetail.OrderName, orderDetail.ServicesId, orderDetail.Deadline.Format("2006-01-02"), orderDetail.Speed, orderDetail.BriefFile, orderDetail.ImportantPoint, orderDetail.AdditionalNotes, orderDetail.OrderSummaryFile)

	if err != nil {
		return entity.OrderDetail{}, errors.New("Error saat membuat order details")
	}

	id, err := execContext.LastInsertId()

	orderDetail.Id = int(id)

	return orderDetail, nil
}

func NewOrderDetailRepository(db *sql.DB) *OrderDetailRepository {
	return &OrderDetailRepository{
		db:        db,
		tableName: "OrderDetail",
	}
}
