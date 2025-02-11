package entity

type Pricing struct {
	Id              int64   `json:"id"`
	BasePrice       float64 `json:"base_price"`
	ServicePrice    float64 `json:"service_price"`
	DeliveryFee     float64 `json:"delivery_fee"`
	Tax             float64 `json:"tax"`
	AdditionalPrice float64 `json:"additional_price"`
	TotalPrice      float64 `json:"total_price"`
	PaymentsId      int64   `json:"payments_id"`
}
