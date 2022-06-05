package services

import (
	productCliente "pan/clients/product"
	"pan/dto"
	"pan/model"
	e "pan/utils/errors"
)

type productService struct{}

//EL SERVICE SE COMUNICA CON EL MODEL (BUSINESS CLASSES) Y CONTROLLER(DTOS)
type productServiceInterface interface {
	GetProductByName(Key string) (dto.ProductsDto, e.ApiError)
	GetProductByCat(key string) (dto.ProductsDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}

}

func (s *productService) GetProductByName(key string) (dto.ProductsDto, e.ApiError) {

	var product []model.Product = productCliente.GetProductByName(key)
	var productDto []dto.ProductDto

	var productAuxDto dto.ProductDto

	if product[0].Id_Product == 0 { // si no se encuentran productos NOT FOUND
		return productDto, e.NewBadRequestApiError("Product Not Found")
	}

	for i := 0; i < len(product); i++ { // guardar datos en productDto[] para devolver al cliente

		productAuxDto.Id_Product = product[i].Id_Product
		productAuxDto.Category = product[i].Category
		productAuxDto.Descripcion = product[i].Desciption
		productAuxDto.Name_product = product[i].Name_product
		productAuxDto.Price = product[i].Price
		productAuxDto.Stock = product[i].Stock
		productDto = append(productDto, productAuxDto)
	}
	return productDto, nil

}

func (s *productService) GetProductByCat(cat string) (dto.ProductsDto, e.ApiError) {

	var product []model.Product = productCliente.GetProductByCat(cat)
	var productDto []dto.ProductDto

	var productAuxDto dto.ProductDto

	if product[0].Id_Product == 0 { // si no se encuentran productos NOT FOUND
		return productDto, e.NewBadRequestApiError("Product Not Found")
	}

	for i := 0; i < len(product); i++ { // guardar datos en productDto[] para devolver al cliente

		productAuxDto.Id_Product = product[i].Id_Product
		productAuxDto.Category = product[i].Category
		productAuxDto.Descripcion = product[i].Desciption
		productAuxDto.Name_product = product[i].Name_product
		productAuxDto.Price = product[i].Price
		productAuxDto.Stock = product[i].Stock
		productDto = append(productDto, productAuxDto)
	}

	return productDto, nil

}
