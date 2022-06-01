package services

import (

	//HAY QUE CAMBIAR ESTOS PATH

	//"log"
	"fmt"
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

	/*

	    PENSAR QUE PASA SI NO ME MANDAN CUALQUIER COSA,
	   	if user.Id == 0 {
	   		return userDto, e.NewBadRequestApiError("user not found")
	   	}
	*/
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

	var product []model.Product = productCliente.GetProductByCat(cat) //LA FUNCION DE PRODUCT CLIENTE NO ESTA BIEN HECHA SE DEBE ESPERAR UN ARRYA DE PRODUCTOS
	var productDto []dto.ProductDto

	var productAuxDto dto.ProductDto

	fmt.Println("client copio bien")

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
