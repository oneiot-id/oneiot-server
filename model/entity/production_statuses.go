package entity

import "time"

type ProductionStatus struct {
	Id               int64     `json:"id"`
	ProductionDate   time.Time `json:"production_date"`
	EstimatedDate    time.Time `json:"estimated_date"`
	LatestStatus     string    `json:"latest_status"`
	ProductionStages string    `json:"production_stages"`
}
