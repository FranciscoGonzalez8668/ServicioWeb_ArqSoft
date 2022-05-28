package product

import (
	"fmt"
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
	var productAux model.Product
	for j := 0; j < 2; j++ {
		fmt.Println("se entra ford client i= ", j)
		productAux.Id_Product = j
		productAux.Category = "electronico"
		productAux.Descripcion = "computadora lenovo"
		productAux.Name_product = "Lenovo Computador"
		productAux.Price = 3.25
		productAux.Stock = 1
		products = append(products, productAux)
	}
	fmt.Println("array created")

	/*Db.Where("Category LIKE ?", cat+"%").Find(&products)

	log.Debug("Products", products)
	*/
	return products

}
