package dto

type CalculatorGetServicesSender struct {
	CityID         string `json:"cityId"`
	ContragentID   string `json:"contragentId"`
	ContragentType string `json:"contragentType"`
	ContractID     string `json:"contractId"`
}
type CalculatorGetServicesReceiver struct {
	CityID         string `json:"cityId"`
	ContragentID   string `json:"contragentId"`
	ContragentType string `json:"contragentType"`
	ContractID     string `json:"contractId"`
}
type CalculatorGetServicesPayer struct {
	ContragentID string `json:"contragentId"`
	ContractID   string `json:"contractId"`
	PayerType    string `json:"payerType"`
}
type CalculatorGetServicesOrderParam struct {
	OrderTypeCode            int    `json:"orderTypeCode"`
	AdditionalOrderTypeCodes []int  `json:"additionalOrderTypeCodes"`
	CashOnDeliveryIndividual bool   `json:"cashOnDeliveryIndividual"`
	IsClientReturn           bool   `json:"isClientReturn"`
	HaveFirstOrderForReturn  bool   `json:"haveFirstOrderForReturn"`
	SellerUUID               string `json:"sellerUuid"`
	OrderCost                int    `json:"orderCost"`
	CalcMode                 string `json:"calcMode,omitempty"`
}
type CalculatorGetServicesPackage struct {
	Length int `json:"length"`
	Width  int `json:"width"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}
type CalculatorGetServicesFilters struct {
	CalcServicesID []string `json:"calcServicesId"`
	CalcModes      []string `json:"calcModes"`
}

type CalculatorGetServicesRequest struct {
	Sender        CalculatorGetServicesSender     `json:"sender"`
	Receiver      CalculatorGetServicesReceiver   `json:"receiver"`
	Payer         CalculatorGetServicesPayer      `json:"payer"`
	OrderParam    CalculatorGetServicesOrderParam `json:"orderParam"`
	InterfaceCode string                          `json:"interfaceCode"`
	CurrencyMark  string                          `json:"currencyMark"`
	CalcDate      string                          `json:"calcDate"`
	Packages      []CalculatorGetServicesPackage  `json:"packages"`
	Filters       CalculatorGetServicesFilters    `json:"filters"`
}
