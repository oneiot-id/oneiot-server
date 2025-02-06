package request

import (
	"oneiot-server/model/entity"
)

type GetOrderRequest struct {
	User  entity.User  `json:"user"`
	Order entity.Order `json:"order"`
}

type GetOrdersRequest struct {
	User entity.User `json:"user"`
}

type CreateOrderRequest struct {
	User        entity.User        `json:"user"`
	OrderDetail entity.OrderDetail `json:"order_detail"`
	Buyer       entity.Buyer       `json:"buyer"`
}

type SetOrderRequest struct {
	User  entity.User  `json:"user"`
	Order entity.Order `json:"order"`
}

type UploadOrderRequest struct {
	User entity.User `json:"user"`
}
