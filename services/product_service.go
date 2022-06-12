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
	GetProductAll() (dto.ProductsDto, e.ApiError)
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
		productAuxDto.Descripcion = product[i].Description
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
		productAuxDto.Descripcion = product[i].Description
		productAuxDto.Name_product = product[i].Name_product
		productAuxDto.Price = product[i].Price
		productAuxDto.Stock = product[i].Stock
		productDto = append(productDto, productAuxDto)
	}

	return productDto, nil

}
func (s *productService) GetProductAll() (dto.ProductsDto, e.ApiError) {
	var productsModel []model.Product = productCliente.GetProductAll()
	var productsDto dto.ProductsDto

	var productsAux dto.ProductDto

	if productsModel[0].Id_Product == 0 {
		return productsDto, e.NewBadRequestApiError("No products in DataBase")
	}

	for k := 0; k < len(productsModel); k++ {
		productsAux.Id_Product = productsModel[k].Id_Product
		productsAux.Category = productsModel[k].Category
		productsAux.Descripcion = productsModel[k].Description
		productsAux.Name_product = productsModel[k].Name_product
		productsAux.Price = productsModel[k].Price
		productsAux.Stock = productsModel[k].Stock
		productsDto = append(productsDto, productsAux)
	}
	return productsDto, nil
}
