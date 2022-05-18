package product

import (
	"pan/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetProductByName(key string) model.Product {

	var product model.Product

	Db.Where("key = ", key).First(&product)
	log.Debug("Product: ", product)

	return product

}

// CORREGIR ESTA FUNCION ESTA MAL
//REVISAR LIBRERIA GORM PARA APRENDER A BUSCAR POR CAT
//:)
//------------------------------------------------------
func GetProductByCat(cat string) model.Products {

	var products model.Products

	Db.Where("cat = ", cat).Find(&products)

	log.Debug("Products", products)

	return products

}

//-------------------------------------------
