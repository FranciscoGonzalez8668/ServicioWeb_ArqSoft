package dto

type Order struct {
	Id_order  int      `json:"id_user"`
	Total     float32  `json:"total"`
	Det_order Detalles `json:"det_orden"`
}
