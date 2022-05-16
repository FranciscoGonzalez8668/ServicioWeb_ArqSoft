package model

type Order struct {
	Id_order    int     `gomr:"primaryKey"`
	Total       float32 `gorm:"type:int;not null"`
	id_detalles int     `gorm:"type:int;not null"` // deberia tener el id de los arreglos de los detalles?? como si fuera una tabal pasarela o como seria emjor
}
