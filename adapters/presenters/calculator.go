package presenters

import "github.com/blackmagiqq/webproxy2/dto"

type CalculatorGetServicesOption struct {
	GeneralServiceID string  `json:"generalServiceId"`
	ServiceName      string  `json:"serviceName"`
	Min              int     `json:"min"`
	Max              int     `json:"max"`
	Price            float32 `json:"price"`
	Ek4id            int     `json:"ek4id"`
}
type CalculatorGetServicesEndpointOption struct {
	Door      []CalculatorGetServicesOption `json:"door"`
	Warehouse []CalculatorGetServicesOption `json:"warehouse"`
	Postamat  []CalculatorGetServicesOption `json:"postamat"`
}
type CalculatorGetServices struct {
	Door      CalculatorGetServicesEndpointOption `json:"door"`
	Warehouse CalculatorGetServicesEndpointOption `json:"warehouse"`
	Postamat  CalculatorGetServicesEndpointOption `json:"postamat,omitempty"`
}

func (p *CalculatorGetServices) FromDTO(d *dto.CalculatorGetServicesResponse) *CalculatorGetServices {
	result := &CalculatorGetServices{}
	for _, s := range d.ServiceList {
		for _, v := range s.ModeDetails {
			var sliceForOption *[]CalculatorGetServicesOption
			switch v.ModeCode {
			case "1":
				sliceForOption = &result.Door.Door
			case "2":
				sliceForOption = &result.Door.Warehouse
			case "3":
				sliceForOption = &result.Warehouse.Door
			case "4":
				sliceForOption = &result.Warehouse.Warehouse
			case "6":
				sliceForOption = &result.Door.Postamat
			case "7":
				sliceForOption = &result.Warehouse.Postamat
			case "8":
				sliceForOption = &result.Postamat.Door
			case "9":
				sliceForOption = &result.Postamat.Warehouse
			case "10":
				sliceForOption = &result.Postamat.Postamat
			default:
				continue
			}

			option := CalculatorGetServicesOption{
				GeneralServiceID: s.GeneralServiceID,
				ServiceName:      s.ServiceName,
				Min:              v.DurationMin,
				Max:              v.DurationMax,
				Price:            v.Price,
				Ek4id:            v.TariffEc4Id,
			}
			*(sliceForOption) = append(*(sliceForOption), option)
		}
	}
	return result
}
