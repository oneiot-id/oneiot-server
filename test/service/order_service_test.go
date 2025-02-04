package service

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"oneiot-server/database"
	"oneiot-server/model/entity"
	"oneiot-server/repository"
	"oneiot-server/service"
	"testing"
	"time"
)

func orderServiceTestBase() service.IOrderService {
	err := godotenv.Load("../../.env")

	if err != nil {
		return nil
	}
	db := database.NewSqlConnection()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, db)

	buyerRepository := repository.NewBuyerRepository(db)

	orderDetailRepository := repository.NewOrderDetailRepository(db)

	orderRepository := repository.NewOrderRepository(db)

	return service.NewOrderService(userService, buyerRepository, orderDetailRepository, orderRepository)
}

func TestCreateNewOrder(t *testing.T) {
	orderService := orderServiceTestBase()

	order := entity.Order{
		IsActive:  false,
		CreatedAt: time.Now(),
	}

	user := entity.User{
		FullName: "testing",
		Email:    "testing@gmail.com",
		Password: "testingpassword",
	}

	orderDetail := entity.OrderDetail{
		OrderName:        "Alat Monitoring Kesehatan Pasien",
		ServicesId:       1,
		Deadline:         time.Now(),
		Speed:            entity.Regular.String(),
		BriefFile:        "alat_monitoring.pdf",
		ImportantPoint:   "",
		AdditionalNotes:  "Dipercepat karena dibutuhkan segera",
		OrderSummaryFile: "order_summary.pdf",
	}

	buyer := entity.Buyer{
		FullName:        "Vincent Kenutama",
		Email:           "",
		PhoneNumber:     "",
		FullAddress:     "",
		AdditionalNotes: "",
	}

	orderDTO, err := orderService.CreateOrder(context.Background(),
		order,
		user,
		orderDetail,
		buyer)

	if err != nil {
		return
	}

	assert.NoError(t, err)
	fmt.Println(orderDTO)
}

func TestGetOrder(t *testing.T) {
	orderService := orderServiceTestBase()

	order := entity.Order{
		Id: 5,
	}

	orderDTO, err := orderService.GetOrderById(context.Background(), order)

	if err != nil {
		return
	}

	assert.NoError(t, err)
	fmt.Println(orderDTO)
}

func TestGetAllUserOrders(t *testing.T) {
	orderService := orderServiceTestBase()

	user := entity.User{
		FullName: "testing",
		Email:    "testing@gmail.com",
		Password: "testingpassword",
	}

	orders, err := orderService.GetAllUserOrder(context.Background(), user)

	if err != nil {
		return
	}

	assert.NoError(t, err)
	fmt.Println(orders)
}
