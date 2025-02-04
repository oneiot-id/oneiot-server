package entity

import "time"

type OrderDetail struct {
	Id               int       `json:"id"`
	OrderName        string    `json:"order_name"`
	ServicesId       int64     `json:"services_id"`
	Deadline         time.Time `json:"deadline"`
	Speed            string    `json:"speed"`
	BriefFile        string    `json:"brief_file"`
	ImportantPoint   string    `json:"important_point"`
	AdditionalNotes  string    `json:"additional_notes"`
	OrderSummaryFile string    `json:"order_summary_file"`
}
