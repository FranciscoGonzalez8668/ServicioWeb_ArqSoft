package model

type Adress struct {
	Id_adress    int    `gorm:"primaryKey"`
	Street_Name  string `gorm:"type:varchar(350);not null"`
	Number       int    `json:"type:int;not null"`
	Neighborhood string `gorm:"type:varchar(350)`
	City         string `gorm:"type:varchar(350);not null`
}

type Adresss []Adress
