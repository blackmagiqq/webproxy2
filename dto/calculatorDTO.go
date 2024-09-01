package dto

type sender struct {
	CityID         string `json:"cityId" binding:"required,uuid"`
	ContragentID   string `json:"contragentId" binding:"required,uuid"`
	ContragentType string `json:"contragentType" binding:"required,oneof=FIZ UR"`
	ContractID     string `json:"contractId" binding:"required,uuid"`
}

type receiver struct {
	CityID         string `json:"cityId" binding:"required,uuid"`
	ContragentID   string `json:"contragentId" binding:"required,uuid"`
	ContragentType string `json:"contragentType" binding:"required,oneof=FIZ UR"`
	ContractID     string `json:"contractId" binding:"required,uuid"`
}

type payer struct {
	ContragentID string `json:"contragentId" binding:"required,uuid"`
	ContractID   string `json:"contractId" binding:"required,uuid"`
	PayerType    string `json:"payerType" binding:"required"`
}

type orderParam struct {
	OrderTypeCode            int    `json:"orderTypeCode" binding:"required,number"`
	AdditionalOrderTypeCodes int    `json:"additionalOrderTypeCodes" binding:"required,number"`
	CashOnDeliveryIndividual bool   `json:"cashOnDeliveryIndividual" binding:"boolean"`
	IsClientReturn           bool   `json:"isClientReturn" binding:"boolean"`
	HaveFirstOrderForReturn  bool   `json:"haveFirstOrderForReturn" binding:"boolean"`
	SellerUUID               string `json:"sellerUuid" binding:"uuid"`
	OrderCost                int    `json:"orderCost" binding:"required,number"`
	CalcMode                 string `json:"calcMode" binding:"required"`
}

type packages struct {
	Length int `json:"length" binding:"required,number"`
	Width  int `json:"width" binding:"required,number"`
	Height int `json:"height" binding:"required,number"`
	Weight int `json:"weight" binding:"required,number"`
}

type filters struct {
	CalcServicesID []string `json:"calcServicesId" binding:"required,dive,required"`
	CalcModes      []string `json:"calcModes" binding:"required,dive,required"`
}

type CalculatorBody struct {
	Sender        sender     `json:"sender"`
	Receiver      receiver   `json:"receiver"`
	Payer         payer      `json:"payer"`
	OrderParam    orderParam `json:"orderParam"`
	InterfaceCode string     `json:"interfaceCode" binding:"required"`
	CurrencyMark  string     `json:"currencyMark" binding:"required"`
	CalcDate      string     `json:"calcDate" binding:"required"`
	Packages      []packages `json:"packages"`
	Filters       filters    `json:"filters"`
}
