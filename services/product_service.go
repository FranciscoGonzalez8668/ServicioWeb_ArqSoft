package services

import (
	"pan/dto"
	"pan/model"
	//HAY QUE CAMBIAR ESTOS PATH
	//	productCliente "mvc-go/clients/user"
	//	"mvc-go/dto"
	//	"mvc-go/model"
)

type productService struct{}

//EL SERVICE SE COMUNICA CON EL MODEL (BUSINESS CLASSES) Y CONTROLLER(DTOS)
type productServiceInterface interface {
	GetProductByKeyCat(key string) dto.ProductDto /*e.Apierror*/
	GetProductByKeyPro(cat string) dto.ProductDto /*e.ApiError*/
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (s *productService) GetProductByKeyPro(key string) dto.ProductDto /*eApiError*/ {

	var product model.Product = productCliente.GetProductByKeyPro(key)
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
	productDto.IdProduct = product.IdProdcut
	productDto.Category = product.Category

	return productDto, nil

}

func (s *productService) GetProductByKeyCat(cat string) dto.ProductsDto /*e.ApiError*/ {

	var products model.products = productCliente.GetProductByKeyCat(cat)
	var usersDto dto.UsersDto

}
