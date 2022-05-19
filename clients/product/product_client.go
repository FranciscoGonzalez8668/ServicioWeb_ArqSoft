package product

import (
	"pan/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetProductByName(key string) model.Product { //HOLA NICO SOY PANCHO CUCHA QUE BUSQUE LA PALABRA IGUAL O PARECIDA?? Y ADEMAS ESO COMPLICA PORQUE ESTAS PASANDO
	// SOLO UN PRODUCTO FIJATE DEBERIAMOS HACER QUE SE COPIEN TODOS LOS PRODUCTOS CON EL NOMBRE ESE
	var product model.Product

	Db.Where("key = ", key).First(&product)
	log.Debug("Product: ", product)

	return product

}

//ACA DEBEMOS PASAR UN ARRAY DE PRODUCTOS , ABRIA QUE VER COMO COPIAR TODOS LOS QUE TIENEN X CATEGORIA
func GetProductByCat(cat string) []model.Product {

	var product []model.Product //DEBERIA SER UN ARRAY? PORQUE CONTIENE MUCHO YA QUE ES <> ????? hay que preguntar esto

	Db.Where("cat <> ?", cat).Find(&product)

	log.Debug("Products", product)

	return product

}
