package entity

import "time"

type DeliveryStatuses struct {
	Id                int64     `json:"id"`
	DeliveryDate      time.Time `json:"delivery_date"`
	ArrivalEstimation time.Time `json:"arrival_estimation"`
	RecipientName     string    `json:"recipient_name"`
	Courier           string    `json:"courier"`
	Address           string    `json:"address"`
	TrackingNumber    string    `json:"tracking_number"`
	DeliveryCourier   string    `json:"delivery_courier"`
}
