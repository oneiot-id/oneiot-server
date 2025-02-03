package entity

import "time"

type OrderDetail struct {
	Id               int64  `json:"id"`
	OrderName        string `json:"order_name"`
	ServicesId       int64
	Deadline         time.Time
	Speed            OrderSpeed
	BriefFile        string
	ImportantPoint   string
	AdditionalNotes  string
	OrderSummaryFile string
}
