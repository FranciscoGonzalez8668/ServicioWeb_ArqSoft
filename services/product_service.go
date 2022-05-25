package services

import (

	//HAY QUE CAMBIAR ESTOS PATH

	//"log"
	productCliente "pan/clients/product"
	"pan/dto"
	"pan/model"

	//	"mvc-go/dto"
	//	"mvc-go/model"
	e "pan/utils/errors"
)

type productService struct{}

//EL SERVICE SE COMUNICA CON EL MODEL (BUSINESS CLASSES) Y CONTROLLER(DTOS)
type productServiceInterface interface {
	GetProductByName(Key string) (dto.ProductDto, e.ApiError)
	GetProductByCat(key string) (dto.ProductsDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (s *productService) GetProductByName(key string) (dto.ProductDto, e.ApiError) { //epi no importado pero en el mvc no esta

	var product model.Product = productCliente.GetProductByName(key)
	var productDto dto.ProductDto

	/*

	    PENSAR QUE PASA SI NO ME MANDAN CUALQUIER COSA,
	   	if user.Id == 0 {
	   		return userDto, e.NewBadRequestApiError("user not found")
	   	}


	*/
	productDto.Name_product = product.Name_product
	productDto.Price = product.Price
	productDto.Stock = product.Stock
	productDto.Id_Product = product.Id_Product
	productDto.Category = product.Category

	return productDto, nil

}

func (s *productService) GetProductByCat(cat string) (dto.ProductsDto, e.ApiError) {

	var product []model.Product = productCliente.GetProductByCat(cat) //LA FUNCION DE PRODUCT CLIENTE NO ESTA BIEN HECHA SE DEBE ESPERAR UN ARRYA DE PRODUCTOS
	var productDto []dto.ProductDto

	var matchProduct int = len(product)

	for i := 0; i < matchProduct; i++ { // guardar datos en productDto[] para devolver al cliente
		productDto[i].Category = product[i].Category
		productDto[i].Id_Product = product[i].Id_Product
		productDto[i].Name_product = product[i].Name_product
		productDto[i].Price = product[i].Price
		productDto[i].Stock = product[i].Stock

	}

	return productDto, nil

}
