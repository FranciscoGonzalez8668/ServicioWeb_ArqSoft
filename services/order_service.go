package services

import (
	/*orderCliente "pan/clients/order"
	 */"pan/dto" /*
		"pan/model"*/

	e "pan/utils/errors"
)

type orderService struct{}

type orderServiceInterface interface {
	OrderHistoy(dto.Token) (dto.OrderDetDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

/*func (s *orderService) OrderHistoy(idUser int) (dto.OrderDetDto, e.ApiError) {
	var orders []model.OrderDet
	orders = orderCliente.GetHistory(idUser)
	var ordersDto dto.OrderDetDto
	if orders[0].Id_order == 0 {
		return ordersDto, e.NewBadRequestApiError("No History")
	}

	for j:=0 ; j<len(orders);j++{
		ordersDto[j]
	}

	return
}*/
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
