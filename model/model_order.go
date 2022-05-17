package model

type Order struct {
	Id_order int     `gomr:"primaryKey"`
	Total    float32 `gorm:"type:decimal(7,6);not null"`
	//id_detalles []int     `gorm:"type:int;not null"` // deberia tener el id de los arreglos de los detalles?? como si fuera una tabla pasarela o como seria mejor
}
