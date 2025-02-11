package general

import (
	"fmt"
	"oneiot-server/model/entity"
	"strconv"
	"testing"
	"time"
)

func TestJsonGenerateProductionType(t *testing.T) {
	var productionStagesData []entity.ProductionStagesData

	for i := 0; i < 10; i++ {
		data := entity.ProductionStagesData{
			Name:   strconv.Itoa(i),
			Date:   time.Now(),
			IsDone: false,
		}

		productionStagesData = append(productionStagesData, data)

	}

	data, err := entity.StringifyProductionStages(productionStagesData)

	if err != nil {
		return
	}

	fmt.Println(data)
}
