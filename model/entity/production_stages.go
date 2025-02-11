package entity

import (
	"encoding/json"
	"time"
)

type ProductionStages struct {
	Data []ProductionStagesData `json:"data"`
}

type ProductionStagesData struct {
	Name   string    `json:"stages_name"`
	Date   time.Time `json:"date"`
	IsDone bool      `json:"is_done"`
}

func StringifyProductionStages(productionData []ProductionStagesData) (string, error) {
	productionStages := ProductionStages{
		Data: productionData,
	}

	marshalled, err := json.Marshal(productionStages)

	if err != nil {
		return "", err
	}

	return string(marshalled), nil
}
