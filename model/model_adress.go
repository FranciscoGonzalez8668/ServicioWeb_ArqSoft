package model

type Adress struct {
	Id_adress    int    `gorm:"primaryKey"`
	Street_Name  string `gorm:"type:varchar(350);not null"`
	Number       int    `gorm:"type:int"`
	Neighborhood string `gorm:"type:varchar(350)"`
	City         string `gorm:"type:varchar(350);not null"`
	Dep          string `gorm:"type:varchar(20)"`
	User         User   `gorm:"type:int;not null"`
	// hace falta poner el id usuaurio ya que se relacionan?
}

type Adresss []Adress
