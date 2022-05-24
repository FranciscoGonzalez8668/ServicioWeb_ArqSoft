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

	Db.Where("key LIKE ? ", "%"+key+"%").First(&product)
	log.Debug("Product: ", product)

	return product

}

//FUNCION LISTA
func GetProductByCat(cat string) []model.Product {

	var products []model.Product

	Db.Where("cat LIKE ?", cat).Find(&products)

	log.Debug("Products", products)

	return products

}
