package dto

type AdressDto struct {
	Street_Name  string  `json:"name"`
	Number       float64 `json:"number"`
	Neighborhood string  `json:"neighbirhood"`
	City         string  `json:"city"`
}

type AdresssDto []AdressDto
