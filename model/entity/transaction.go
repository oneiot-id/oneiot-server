package entity

import "time"

type Transaction struct {
	Id                   int64     `json:"id"`
	UserId               int64     `json:"user_id"`
	OrderId              int64     `json:"order_id"`
	PricingId            int64     `json:"pricing_id"`
	ProductionStatusesId int64     `json:"production_statuses_id"`
	DeliveryStatusesId   int64     `json:"delivery_statuses_id"`
	Status               string    `json:"status"`
	CreatedAt            time.Time `json:"created_at"`
	Complained           bool      `json:"complained"`
}
