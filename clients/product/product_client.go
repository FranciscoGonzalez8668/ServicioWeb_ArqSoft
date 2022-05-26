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

	var products = make([]model.Product, 1)
	fmt.Println("array length: ", len(products))
	for j := 0; j < 2; j++ {
		fmt.Println("se entra ford client i= ", j)
		products[j].Id_Product = j
		products[j].Category = "electronico"
		products[j].Descripcion = "computadora lenovo"
		products[j].Name_product = "Lenovo Computador"
		products[j].Price = 3.25
		products[j].Stock = 1
	}
	fmt.Println("array created")

	/*Db.Where("cat LIKE ?", cat+"%").Find(&products)

	log.Debug("Products", products)
	*/
	return products

}
