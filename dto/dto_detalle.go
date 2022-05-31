package dto

type DetalleDto struct {
	Product   ProductDto `json:"product"`
	Price_det float32    `json:"price_det"`
	Cant_det  int        `json:"cant_det"`
}

type DetallesDto struct {
	Detalles_id_usuario string       `json:"id_usuario"`
	DetalleArray        []DetalleDto `json:"detalles"`
}

type DetalleDetDto struct {
	Detalle_id_product string  `json:"id_product"`
	Price_det          float32 `json:"price_det"`
	Cant_det           int     `json:"cant_det"`
	Product_Name       string  `json:"product_name"`
}

type DetallesDetDto []DetalleDetDto
