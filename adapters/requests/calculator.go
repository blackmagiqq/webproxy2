package requests

type CalculatorGetServices struct {
	Sender struct {
		CityID         string `json:"cityId" binding:"required,uuid" faker:"uuid_hyphenated"`
		ContragentID   string `json:"contragentId" binding:"omitempty,uuid" faker:"uuid_hyphenated"`
		ContragentType string `json:"contragentType" binding:"required,oneof=FIZ UR" faker:"oneof:FIZ, UR"`
		ContractID     string `json:"contractId" binding:"omitempty,uuid" faker:"uuid_hyphenated"`
	} `json:"sender" binding:"required"`
	Receiver struct {
		CityID         string `json:"cityId" binding:"required,uuid" faker:"uuid_hyphenated"`
		ContragentID   string `json:"contragentId" binding:"omitempty,uuid" faker:"uuid_hyphenated"`
		ContragentType string `json:"contragentType" binding:"required,oneof=FIZ UR" faker:"oneof:FIZ, UR"`
		ContractID     string `json:"contractId" binding:"omitempty,uuid" faker:"uuid_hyphenated"`
	} `json:"receiver" binding:"required"`
	Payer struct {
		ContragentID string `json:"contragentId" binding:"omitempty,uuid" faker:"uuid_hyphenated"`
		ContractID   string `json:"contractId" binding:"omitempty,uuid" faker:"uuid_hyphenated"`
		PayerType    string `json:"payerType" binding:"required,oneof=sender receiver" faker:"oneof:sender, receiver"`
	} `json:"payer" binding:"required"`
	OrderParam struct {
		OrderTypeCode            int    `json:"orderTypeCode" binding:"required,number"`
		AdditionalOrderTypeCodes []int  `json:"additionalOrderTypeCodes" binding:"required,dive,omitempty,number"`
		CashOnDeliveryIndividual bool   `json:"cashOnDeliveryIndividual" binding:"boolean"`
		IsClientReturn           bool   `json:"isClientReturn" binding:"boolean"`
		HaveFirstOrderForReturn  bool   `json:"haveFirstOrderForReturn" binding:"boolean"`
		SellerUUID               string `json:"sellerUuid" binding:"omitempty,uuid" faker:"uuid_hyphenated"`
		OrderCost                int    `json:"orderCost" binding:"omitempty,number"`
		CalcMode                 string `json:"calcMode" binding:"omitempty"`
	} `json:"orderParam"  binding:"required"`
	InterfaceCode string `json:"interfaceCode" binding:"required"`
	CurrencyMark  string `json:"currencyMark" binding:"required"`
	CalcDate      string `json:"calcDate"`
	Packages      []struct {
		Length int `json:"length" binding:"required,number"`
		Width  int `json:"width" binding:"required,number"`
		Height int `json:"height" binding:"required,number"`
		Weight int `json:"weight" binding:"required,number"`
	} `json:"packages" binding:"required"`
	Filters struct {
		CalcServicesID []string `json:"calcServicesId" binding:"omitempty,dive,uuid"`
		CalcModes      []string `json:"calcModes" binding:"omitempty"`
	} `json:"filters" binding:"required"`
	Language string `json:"language" binding:"omitempty"`
}
