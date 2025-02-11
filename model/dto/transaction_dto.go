package dto

import "oneiot-server/model/entity"

type TransactionDto struct {
	Transaction      entity.Transaction      `json:"transaction"`
	Pricing          entity.Pricing          `json:"pricing"`
	Payment          entity.Payment          `json:"payment"`
	ProductionStatus entity.ProductionStatus `json:"production_status"`
	DeliveryStatus   entity.DeliveryStatuses `json:"delivery_status"`
}
