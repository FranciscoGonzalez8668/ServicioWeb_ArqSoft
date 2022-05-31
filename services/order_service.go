package services

import (
	orderCliente "pan/clients/order"
	"pan/dto"
	"pan/model"

	e "pan/utils/errors"
)

type orderService struct{}

type orderServiceInterface interface {
	OrderHistoy(int) ([]dto.OrderDetDto, e.ApiError)
	NewOrder(detallesDto []dto.DetalleDto) (dto.OrderDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) OrderHistoy(idUser int) ([]dto.OrderDetDto, e.ApiError) {
	var orders []model.OrderDet
	var id_order []int
	var DetalleAux dto.DetalleDetDto
	orders, id_order = orderCliente.GetHistory(idUser)
	var ordersDto []dto.OrderDetDto
	for j := 0; j < len(id_order); j++ {
		ordersDto[j].Id_order = id_order[j]
		for i := 0; i < len(orders); i++ {
			if orders[i].Id_order == ordersDto[j].Id_order {
				DetalleAux.Cant_det = orders[i].Cantidad
				DetalleAux.Price_det = orders[i].Precio_Moment
				DetalleAux.Product_Name = orders[i].Product
				ordersDto[j].Det_order = append(ordersDto[j].Det_order)
			}
		}
		for k := 0; k < len(orders); k++ {
			if orders[k].Id_order == ordersDto[j].Id_order {
				ordersDto[j].Total = orders[k].Total
				break
			}
		}

	}
	return ordersDto, nil
}
func (s *orderService) NewOrder(detallesDto []dto.DetalleDto) (dto.OrderDto, e.ApiError) {
	var order model.Order
	var detalles []model.DetalleDet
	var total float32 = 0
	for j := 0; j < len(detalles); j++ {
		detalles[j].Cantidad = detallesDto[j].Cant_det
		detalles[j].Id_Product = detallesDto[j].Detalle_id_product
		detalles[j].Precio_Moment = detallesDto[j].Price_det
		total = total + detallesDto[j].Price_det
	}
	order.Total = total
	order = orderCliente.NewOrder(order, detalles)
	//hay una diferencia entre los dto y los modelos ya que los dto tienen datos de otras tablas de la db
	var orderDto dto.OrderDto
	orderDto.Id_order = order.Id_order
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
