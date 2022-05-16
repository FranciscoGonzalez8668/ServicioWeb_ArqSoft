package dto

type DetalleDto struct {
	Name_product string  `json:"name_product"`
	Price_det    float32 `json:"price_det"`
	Cant_det     int     `json:"cant_det"`
}

type DetallesDto []DetalleDto
