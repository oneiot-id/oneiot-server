package entity

import "time"

type Order struct {
	Id                 int64     `json:"id"`
	UserId             int64     `json:"user_id"`
	BuyerId            int64     `json:"buyer_id"`
	OrderDetailId      int64     `json:"order_detail_id"`
	IsActive           bool      `json:"is_active"`
	Confirmed          bool      `json:"confirmed"`
	CreatedAt          time.Time `json:"created_at"`
	TransactionCreated bool      `json:"transaction_created"`
}

type OrderDTO struct {
	Order       Order       `json:"order"`
	Buyer       Buyer       `json:"buyer"`
	OrderDetail OrderDetail `json:"order_detail"`
}
