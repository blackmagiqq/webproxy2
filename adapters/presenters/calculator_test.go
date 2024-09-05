package presenters

import (
	"testing"

	"github.com/blackmagiqq/webproxy2/dto"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func TestMakeViewFromDTO(t *testing.T) {
	var service dto.CalculatorGetServicesResponseService
	faker.FakeData(&service)

	var modeDoorDoor dto.CalculatorGetServicesResponseModeDetail
	faker.FakeData(&modeDoorDoor)
	modeDoorDoor.ModeCode = "1"

	var modeDoorWarehouse dto.CalculatorGetServicesResponseModeDetail
	faker.FakeData(&modeDoorWarehouse)
	modeDoorWarehouse.ModeCode = "2"

	var modeWarehouseDoor dto.CalculatorGetServicesResponseModeDetail
	faker.FakeData(&modeWarehouseDoor)
	modeWarehouseDoor.ModeCode = "3"

	service.ModeDetails = []dto.CalculatorGetServicesResponseModeDetail{modeDoorDoor, modeDoorWarehouse, modeWarehouseDoor}

	responseDTO := dto.CalculatorGetServicesResponse{ServiceList: []dto.CalculatorGetServicesResponseService{service}}

	presenter := &CalculatorGetServices{}
	view := presenter.FromDTO(&responseDTO)

	require.Equal(t, view.Door.Door[0].Ek4id, modeDoorDoor.TariffEc4Id)
	require.Equal(t, view.Door.Warehouse[0].Ek4id, modeDoorWarehouse.TariffEc4Id)
	require.Equal(t, view.Warehouse.Door[0].Ek4id, modeWarehouseDoor.TariffEc4Id)
}
