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
	GetProductByCat(key string) (dto.ProductDto, e.ApiError, int)
	GetProductByName(Key string) (dto.ProductDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{} //no entiendo porque no funciona
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

func (s *productService) GetProductByCat(cat string) (dto.ProductsDto, e.ApiError, int) {

	var products model.Products = productCliente.GetProductByCat(cat) //LA FUNCION DE PRODUCT CLIENTE NO ESTA BIEN HECHA SE DEBE ESPERAR UN ARRYA DE PRODUCTOS
	var productsDto dto.ProductsDto

	var matchProduct int = len(products)

	for i := 0; i < matchProduct; i++ {
		productsDto[i].Category = products[i].Category
		productsDto[i].Id_Product = products[i].Id_Product
		productsDto[i].Name_product = products[i].Name_product
		productsDto[i].Price = products[i].Price
		productsDto[i].Stock = products[i].Stock

	}

	return productsDto, nil, matchProduct

}
