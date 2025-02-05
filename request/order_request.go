package request

import (
	"oneiot-server/model/entity"
)

type GetOrderRequest struct {
	User  entity.User  `json:"user_pictures"`
	Order entity.Order `json:"order"`
}

type GetOrdersRequest struct {
	User entity.User `json:"user_pictures"`
}

type CreateOrderRequest struct {
	User        entity.User        `json:"user_pictures"`
	OrderDetail entity.OrderDetail `json:"order_detail"`
	Buyer       entity.Buyer       `json:"buyer"`
}
