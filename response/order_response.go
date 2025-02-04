package response

import "oneiot-server/model/entity"

type GetAllOrdersResponse struct {
	Orders []entity.OrderDTO `json:"orders"`
}

type CreateOrderResponse struct {
	Order entity.OrderDTO `json:"order"`
}
