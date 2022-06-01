package model

type Order struct {
	Id_Orden int     `gomr:"primaryKey"`
	Id_User  int     `gorm:"type:int"`
	Total    float32 `gorm:"type:float"`
	//id_detalles []int     `gorm:"type:int;not null"` // deberia tener el id de los arreglos de los detalles?? como si fuera una tabla pasarela o como seria mejor
}

/*type OrderDet struct {
	Id_Orden      int
	Total         float32
	Product       string
	Cantidad      int
	Precio_Moment float32
}*/
