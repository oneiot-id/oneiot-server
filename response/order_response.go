package response

import "oneiot-server/model/entity"

type GetAllOrdersResponse struct {
	Orders []entity.OrderDTO `json:"orders"`
}

type CreateOrderResponse struct {
	Order entity.OrderDTO `json:"order"`
}

type UpdateOrderResponse struct {
	User  entity.User  `json:"user"`
	Order entity.Order `json:"order"`
}

type UpdateBriefFile struct {
	User     entity.User     `json:"user"`
	OrderDTO entity.OrderDTO `json:"order_dto"`
}
