package entity

import "time"

type OrderDetail struct {
	Id               int    `json:"id"`
	OrderName        string `json:"order_name"`
	ServicesId       int64
	Deadline         time.Time
	Speed            string
	BriefFile        string
	ImportantPoint   string
	AdditionalNotes  string
	OrderSummaryFile string
}
