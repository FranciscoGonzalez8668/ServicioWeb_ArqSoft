package model

type Detalle struct {
	Id_detalle    int     `gorm:"type:primaryKey"`
	Cantidad      int     `gorm:"type:int;not null"`
	Precio_Moment float32 `gorm:"type:decimal(7,6);not null"`
}
