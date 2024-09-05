package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/blackmagiqq/webproxy2/adapters/presenters"
	"github.com/blackmagiqq/webproxy2/adapters/requests"
	"github.com/blackmagiqq/webproxy2/dto"
	"github.com/bxcodec/faker/v4"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	godotenv.Load("../../.env")

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

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/calculator/getServices" {
			responseJSON, _ := json.Marshal(responseDTO)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJSON)
			return
		}

		http.NotFound(w, r)
	}))
	defer mockServer.Close()

	originalHost := os.Getenv("CALCULATOR_HOST")
	t.Setenv("CALCULATOR_HOST", mockServer.URL)
	defer func() { t.Setenv("CALCULATOR_HOST", originalHost) }()

	router := SetupRouter()

	w := httptest.NewRecorder()

	var payload requests.CalculatorGetServices
	faker.FakeData(&payload)
	payload.Filters.CalcServicesID = []string{faker.UUIDHyphenated()}
	payloadJSON, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/v2/calculator/getServices", strings.NewReader(string(payloadJSON)))
	router.ServeHTTP(w, req)

	expected := presenters.CalculatorGetServices{
		Door: presenters.CalculatorGetServicesEndpointOption{
			Door: []presenters.CalculatorGetServicesOption{
				{
					Ek4id:            modeDoorDoor.TariffEc4Id,
					GeneralServiceID: service.GeneralServiceID,
					Max:              modeDoorDoor.DurationMax,
					Min:              modeDoorDoor.DurationMin,
					Price:            modeDoorDoor.Price,
					ServiceName:      service.ServiceName,
				},
			},
			Warehouse: []presenters.CalculatorGetServicesOption{
				{
					Ek4id:            modeDoorWarehouse.TariffEc4Id,
					GeneralServiceID: service.GeneralServiceID,
					Max:              modeDoorWarehouse.DurationMax,
					Min:              modeDoorWarehouse.DurationMin,
					Price:            modeDoorWarehouse.Price,
					ServiceName:      service.ServiceName,
				},
			},
		},
		Warehouse: presenters.CalculatorGetServicesEndpointOption{
			Door: []presenters.CalculatorGetServicesOption{
				{
					Ek4id:            modeWarehouseDoor.TariffEc4Id,
					GeneralServiceID: service.GeneralServiceID,
					Max:              modeWarehouseDoor.DurationMax,
					Min:              modeWarehouseDoor.DurationMin,
					Price:            modeWarehouseDoor.Price,
					ServiceName:      service.ServiceName,
				},
			},
		},
	}
	expectedJSON, _ := json.Marshal(expected)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, string(expectedJSON), w.Body.String())
}
