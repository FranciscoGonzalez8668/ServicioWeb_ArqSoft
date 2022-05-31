package dto

type OrderDto struct {
	Id_order  int          `json:"id_user"`
	Total     float32      `json:"total"`
	Det_order []DetalleDto `json:"det_orden"`
	Adress    AdressDto    `json:"adress"`
}

type OrdersHistoryDto struct {
	Id_User int        `json:"id_user"`
	Order   []OrderDto `json:"order"`
}

type OrderDetDto struct {
	Id_order  int            `json:"id_user"`
	Total     float32        `json:"total"`
	Det_order DetallesDetDto `json:"det_orden"`
}
