package dto

type CalculatorGetServicesResponseModeDetail struct {
	ModeCode                  string  `json:"modeCode" faker:"oneof: 1, 2, 3, 4, 6, 7, 8, 9, 10"`
	ModeName                  string  `json:"modeName" faker:"word"`
	DurationMin               int     `json:"durationMin"`
	DurationMax               int     `json:"durationMax"`
	CalendarPeriodMin         int8    `json:"calendarPeriodMin"`
	CalendarPeriodMax         int8    `json:"calendarPeriodMax"`
	TimeCourierInvitation     string  `json:"timeCourierInvitation" faker:"word"`
	TimeWarehouse             string  `json:"timeWarehouse"`
	TimePvz                   string  `json:"timePvz"`
	Price                     float32 `json:"price"`
	TariffEc4Id               int     `json:"tariffEc4Id"`
	ForecastDeliveryDateRange struct {
		Min string `json:"min"`
		Max string `json:"max"`
	} `json:"forecastDeliveryDateRange"`
}

type CalculatorGetServicesResponseService struct {
	GeneralServiceID   string                                    `json:"generalServiceId" faker:"uuid_hyphenated"`
	ServiceName        string                                    `json:"serviceName" faker:"word"`
	ServiceDescription string                                    `json:"serviceDescription" faker:"word"`
	SpeedTypeCode      string                                    `json:"speedTypeCode"`
	ModeDetails        []CalculatorGetServicesResponseModeDetail `json:"modeDetails"`
	CalcWeight         float32                                   `json:"calcWeight"`
}

type CalculatorGetServicesResponse struct {
	ServiceList []CalculatorGetServicesResponseService `json:"serviceList"`
}
