package entity

import "time"

type Transaction struct {
	Id                   int       `json:"id"`
	UserId               int       `json:"user_id"`
	OrderId              int       `json:"order_id"`
	PricingId            int       `json:"pricing_id"`
	ProductionStatusesId int       `json:"production_statuses_id"`
	DeliveryStatusesId   int       `json:"delivery_statuses_id"`
	Status               string    `json:"status"`
	CreatedAt            time.Time `json:"created_at"`
	Complained           bool      `json:"complained"`
}
