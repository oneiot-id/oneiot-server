package service

import (
	"context"
	"errors"
	"oneiot-server/model/entity"
	"oneiot-server/repository"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, order entity.Order, user entity.User, detail entity.OrderDetail, buyer entity.Buyer) (entity.OrderDTO, error)
	GetAllUserOrder(ctx context.Context, user entity.User) ([]entity.OrderDTO, error)
	GetOrderById(ctx context.Context, order entity.Order) (entity.OrderDTO, error)
	SetStatus(ctx context.Context, order entity.Order) (entity.OrderDTO, error)
	UploadBriefFile(ctx context.Context, orderDTO entity.OrderDTO, checkIdExisted bool) (entity.OrderDTO, error)
}

type OrderService struct {
	userService           IUserService
	buyerRepository       repository.IBuyerRepository
	orderDetailRepository repository.IOrderDetailRepository
	orderRepository       repository.IOrderRepository
}

func (service *OrderService) UploadBriefFile(ctx context.Context, orderDTO entity.OrderDTO, checkId bool) (entity.OrderDTO, error) {
	//Cek apakah order dengan id ini ada

	var order entity.Order

	//Skip jika checkId tidak digunakan
	if checkId {
		o, err := service.orderRepository.GetOrderById(ctx, orderDTO.Order.Id)

		if err != nil {
			return entity.OrderDTO{}, err
		}

		order = o
	}

	//Jika ada insert
	orderDetail, err := service.orderDetailRepository.UpdateBriefFile(ctx, orderDTO.OrderDetail)

	if err != nil {
		return entity.OrderDTO{}, err
	}

	return entity.OrderDTO{
		Order:       order,
		Buyer:       orderDTO.Buyer,
		OrderDetail: orderDetail,
	}, nil
}

func (service *OrderService) SetStatus(ctx context.Context, order entity.Order) (entity.OrderDTO, error) {

	updatedOrder, err := service.orderRepository.SetOrderStatus(ctx, order)

	if err != nil {
		return entity.OrderDTO{}, err
	}

	return entity.OrderDTO{Order: updatedOrder}, nil
}

func (service *OrderService) GetOrderById(ctx context.Context, order entity.Order) (entity.OrderDTO, error) {

	order, err := service.orderRepository.GetOrderById(ctx, order.Id)

	buyer, err := service.buyerRepository.GetById(ctx, entity.Buyer{Id: order.BuyerId})

	orderDetail, err := service.orderDetailRepository.GetOrderById(ctx, entity.OrderDetail{Id: int(order.OrderDetailId)})

	if err != nil {
		return entity.OrderDTO{}, err
	}

	return entity.OrderDTO{
		Order:       order,
		Buyer:       buyer,
		OrderDetail: orderDetail,
	}, nil
}

func (service *OrderService) GetAllUserOrder(ctx context.Context, user entity.User) ([]entity.OrderDTO, error) {

	user, err := service.userService.GetUserByID(ctx, user.Id)

	if err != nil {
		return nil, err
	}

	// 2. we get all the order
	order, err := service.orderRepository.GetOrdersByUserId(ctx, user)

	if err != nil {
		return nil, err
	}

	//ToDo: Jika ingin membuat slice untuk order dto mungkin dapat menggunakan

	var orders []entity.OrderDTO

	for _, o := range order {
		// Loop untuk order lalu gunakan untuk membuat order dto
		var orderDTO entity.OrderDTO

		buyer, err := service.buyerRepository.GetById(ctx, entity.Buyer{Id: o.BuyerId})
		orderDetail, err := service.orderDetailRepository.GetOrderById(ctx, entity.OrderDetail{Id: int(o.OrderDetailId)})

		if err != nil {
			return nil, err
		}

		orderDTO.Order = o
		orderDTO.OrderDetail = orderDetail
		orderDTO.Buyer = buyer

		orders = append(orders, orderDTO)
	}

	return orders, nil
}

func (service *OrderService) CreateOrder(ctx context.Context, order entity.Order, user entity.User, detail entity.OrderDetail, buyer entity.Buyer) (entity.OrderDTO, error) {

	// [x] 1. First we have to get the user_pictures who creating it is it exist or not
	user, err := service.userService.GetUserByID(ctx, user.Id)

	if err != nil {
		return entity.OrderDTO{}, errors.New("tidak dapat membuat order, user tidak ditemukan")
	}

	order.UserId = int64(user.Id)

	// 2. Then we can create the buyer
	buyer, err = service.buyerRepository.Create(ctx, buyer)

	if err != nil {
		return entity.OrderDTO{}, err
	}

	order.BuyerId = int64(buyer.Id)

	// 3. Then we can create the order details
	orderDetail, err := service.orderDetailRepository.CreateOrderDetail(ctx, detail)

	if err != nil {
		return entity.OrderDTO{}, err
	}

	order.OrderDetailId = int64(orderDetail.Id)

	// 4. Last we create the order
	createOrder, err := service.orderRepository.CreateOrder(ctx, order)

	if err != nil {
		return entity.OrderDTO{}, err
	}

	//To get the DTOs we have to gather all the information as one (anjay)

	return entity.OrderDTO{
		Order:       createOrder,
		Buyer:       buyer,
		OrderDetail: orderDetail,
	}, nil
}

func NewOrderService(userService IUserService, buyerRepository repository.IBuyerRepository, orderDetailRepository repository.IOrderDetailRepository, orderRepository repository.IOrderRepository) IOrderService {
	return &OrderService{
		userService:           userService,
		buyerRepository:       buyerRepository,
		orderDetailRepository: orderDetailRepository,
		orderRepository:       orderRepository,
	}
}
