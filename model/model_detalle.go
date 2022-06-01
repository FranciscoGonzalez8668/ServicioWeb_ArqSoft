package model

type Detalle struct {
	Id_Orden      int     `gorm:"type:primaryKey"`
	Id_Detalle    int     `gorm:"type:primaryKey"`
	Id_Product    int     `gorm:"type:int"`
	Cantidad      int     `gorm:"type:int;not null"`
	Precio_Moment float32 `gorm:"type:float;not null"`
}

/*type DetalleDet struct {
	//Id_detalle    int     `gorm"type:primaryKey"`
	Cantidad      int     `gorm:"type:int;not null"`
	Precio_Moment float32 `gorm:"type: decimal (7,6)"`
	Id_Product    int
	Id_Order      int
}*/
