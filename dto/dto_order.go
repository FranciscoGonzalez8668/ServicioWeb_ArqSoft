package dto

type OrderDto struct {
	Id_order  int         `json:"id_user"`
	Total     float32     `json:"total"`
	Det_order DetallesDto `json:"det_orden"`
	Adress    AdressDto   `json:"adress"`
}