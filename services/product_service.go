package services

import (

	//HAY QUE CAMBIAR ESTOS PATH
	"pan/dto"
	"pan/model"
	productCliente "servicioweb_arqsoft/clients/product"
	//	"mvc-go/dto"
	//	"mvc-go/model"
	e "servicioweb_arqsoft/utils/errors"
)

type productService struct{}

//EL SERVICE SE COMUNICA CON EL MODEL (BUSINESS CLASSES) Y CONTROLLER(DTOS)
type productServiceInterface interface {
	GetProductByKeyCat(key string) dto.ProductDto /*e.Apierror*/
	GetProductByKeyPro(Key string) dto.ProductDto /*e.ApiError*/
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
	productDto.Cost = product.Cost
	productDto.Stock = product.Stock
	productDto.Id_Product = product.Id_Product
	productDto.Category = product.Category

	return productDto, nil

}

func (s *productService) GetProductByCat(cat string) dto.ProductsDto /*e.ApiError*/ {

	var products model.Products = productCliente.GetProductByCat(cat)
	var usersDto dto.UsersDto

	return products

}
