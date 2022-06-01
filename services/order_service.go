package services

import (
	orderCliente "pan/clients/order"
	productClient "pan/clients/product"
	"pan/dto"
	"pan/model"
	e "pan/utils/errors"
)

type orderService struct{}

type orderServiceInterface interface {
	OrderHistoy(int) (dto.OrdersHistoryDto, e.ApiError)
	NewOrder(detallesDto []dto.DetalleDto) (dto.OrderDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) OrderHistoy(idUser int) (dto.OrdersHistoryDto, e.ApiError) {
	var orders []model.Order
	var detalles []model.Detalle
	var product model.Product
	var order dto.OrdersHistoryDto
	orders = orderCliente.GetOrders(idUser)
	for j := 0; j < len(orders); j++ {
		detalles = orderCliente.GetDetalles(orders[j].Id_Orden)
		for k := 0; k < len(detalles); k++ {
			product = productClient.GetProductById(detalles[k].Id_Product)
			order.Order[j].Det_order[k].Product.Id_Product = product.Id_Product
			order.Order[j].Det_order[k].Product.Name_product = product.Name_product
			order.Order[j].Det_order[k].Product.Category = product.Category
			order.Order[j].Det_order[k].Product.Price = product.Price
			order.Order[j].Det_order[k].Product.Stock = product.Stock
			order.Order[j].Det_order[k].Product.Descripcion = product.Desciption

			order.Order[j].Det_order[k].Cant_det = detalles[k].Cantidad
			order.Order[j].Det_order[k].Price_det = detalles[k].Precio_Moment
		}
	}
	order.Id_User = idUser
	return order, nil

}
func (s *orderService) NewOrder(detallesDto []dto.DetalleDto) (dto.OrderDto, e.ApiError) {
	var order model.Order
	var detalles []model.Detalle
	var total float32 = 0

	//falta escribir todos los datos en el producto para mostrar bien la order
	for j := 0; j < len(detalles); j++ {
		detalles[j].Cantidad = detallesDto[j].Cant_det
		detalles[j].Id_Product = detallesDto[j].Product.Id_Product
		detalles[j].Precio_Moment = detallesDto[j].Price_det
		total = total + detallesDto[j].Price_det
	}
	order.Total = total
	order = orderCliente.NewOrder(order, detalles)
	var orderDto dto.OrderDto
	if order.Id_Orden == 0 {
		return orderDto, e.NewBadRequestApiError("No stock enought")
	}
	//hay una diferencia entre los dto y los modelos ya que los dto tienen datos de otras tablas de la db

	orderDto.Id_order = order.Id_Orden
	orderDto.Total = order.Total

	return orderDto, nil
}

/*
 ACA VA MI INTENTO
func (s *orderService) OrderHistoy(idUser int) (dto.OrderDetDto, e.ApiError){


	var orders []model.OrderDet
	orders = orderCliente.GetHistory (idUser)
	var orders dto.OrderDetDto

	for(i=0, i</* cantidad de orderdenes , i++){

		ordersDto[i] = [i]model.OrderDet



	}

	return ordersDto



}


*/
