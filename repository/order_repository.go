package repository

import (
	"context"
	"database/sql"
	"errors"
	"oneiot-server/helper"
	"oneiot-server/model/entity"
	"time"
)

type IOrderRepository interface {
	CreateOrder(ctx context.Context, order entity.Order) (entity.Order, error)

	//I don't know if the user_pictures is be able to delete the order, but we might try it later for this
	DeleteOrderById(ctx context.Context, orderId int64) error

	GetOrderById(ctx context.Context, orderId int64) (entity.Order, error)
	GetOrdersByUserId(ctx context.Context, user entity.User) ([]entity.Order, error)
	SetOrderStatus(ctx context.Context, order entity.Order) (entity.Order, error)
}

type OrderRepository struct {
	db *sql.DB
}

func (repository *OrderRepository) CreateOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	query := "INSERT INTO Orders(UserId, BuyerId, OrderDetailId, IsActive, CreatedAt)" +
		"VALUES(?, ?, ?, ?, ?)"

	execContext, err := repository.db.ExecContext(ctx, query, order.UserId, order.BuyerId, order.OrderDetailId, order.IsActive, order.CreatedAt.Format("2006-01-02 15:04:05"))

	if err != nil {
		return entity.Order{}, errors.New("Error saat membuat order")
	}

	order.Id, err = execContext.LastInsertId()

	return order, nil
}

func (repository *OrderRepository) DeleteOrderById(ctx context.Context, orderId int64) error {

	query := "DELETE FROM Orders WHERE Id=?"

	_, err := repository.db.ExecContext(ctx, query, orderId)

	if err != nil {
		return errors.New("Error saat menghapus order")
	}

	return nil
}

func (repository *OrderRepository) GetOrderById(ctx context.Context, orderId int64) (entity.Order, error) {
	query := "SELECT Id, UserId, BuyerId, OrderDetailId, IsActive, CreatedAt, Confirmed From Orders WHERE Id=?"

	row, err := repository.db.QueryContext(ctx, query, orderId)

	if err != nil {
		return entity.Order{}, errors.New("Error saat mendapatkan order dengan id " + string(rune(orderId)))
	}

	if !row.Next() {
		return entity.Order{}, errors.New("Error tidak ditemukan order dengan id " + string(rune(orderId)))
	}

	var order entity.Order
	var createdAt string

	err = row.Scan(&order.Id, &order.UserId, &order.BuyerId, &order.OrderDetailId, &order.IsActive, &createdAt, &order.Confirmed)

	if err != nil {
		return entity.Order{}, errors.New("Error saat scannning order dengan id " + string(rune(orderId)))
	}

	order.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)

	return order, nil
}

func (repository *OrderRepository) GetOrdersByUserId(ctx context.Context, user entity.User) ([]entity.Order, error) {
	query := "SELECT Id, UserId, BuyerId, OrderDetailId, IsActive, CreatedAt, Confirmed From Orders WHERE UserId = ? "

	rows, err := repository.db.QueryContext(ctx, query, user.Id)

	if err != nil {
		return nil, errors.New("Error saat melakukan query ke tabel order")
	}

	var orders []entity.Order

	for rows.Next() {
		var order entity.Order
		var createdAt string

		err = rows.Scan(&order.Id, &order.UserId, &order.BuyerId, &order.OrderDetailId, &order.IsActive, &createdAt, &order.Confirmed)

		order.CreatedAt = helper.StringToDateTime(createdAt)
		orders = append(orders, order)
	}

	if len(orders) == 0 {
		return nil, errors.New("User belum memesan sesuatu")
	}

	return orders, nil
}

func (repository *OrderRepository) SetOrderStatus(ctx context.Context, order entity.Order) (entity.Order, error) {
	query := "UPDATE orders SET IsActive = ?, Confirmed = ? " +
		"Where Id = ?"

	_, err := repository.db.ExecContext(ctx, query, order.IsActive, order.Confirmed, order.Id)

	if err != nil {
		return entity.Order{}, errors.New("Error saat mengupdate status order dengan id" + string(rune(order.Id)))
	}

	return order, nil
}

func NewOrderRepository(db *sql.DB) IOrderRepository {
	return &OrderRepository{
		db: db,
	}
}
