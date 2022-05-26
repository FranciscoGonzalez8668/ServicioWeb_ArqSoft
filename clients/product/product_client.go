package product

import (
	"pan/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

//FUNCION LISTA

func GetProductByName(key string) model.Product {
	var product model.Product

	Db.Where("key LIKE ? ", "%"+key+"%").First(&product)
	log.Debug("Product: ", product)

	return product
}

//FUNCION LISTA
//
//Params: 	string
//Return:	array model.producto
//Proces:	Busca en BD productos con categoria @params@
//
//
func GetProductByCat(cat string) []model.Product { //

	var products []model.Product

	Db.Where("cat LIKE ?", cat+"%").Find(&products)

	log.Debug("Products", products)

	return products

}
