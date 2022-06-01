package model

type Detalle struct {
	Id_Order      int     `gorm:"type:primaryKey"`
	Id_detalle    int     `gorm:"type:primaryKey"`
	Id_Product    int     `gorm:"type:int;not null"`
	Cantidad      int     `gorm:"type:int;not null"`
	Precio_Moment float32 `gorm:"type:decimal(7,6);not null"`
}

type DetalleDet struct {
	//Id_detalle    int     `gorm"type:primaryKey"`
	Cantidad      int     `gorm:"type:int;not null"`
	Precio_Moment float32 `gorm:"type: decimal (7,6)"`
	Id_Product    int
	Id_Order      int
}
