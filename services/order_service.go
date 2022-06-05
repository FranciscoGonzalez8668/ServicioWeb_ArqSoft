package services

import (
	orderCliente "pan/clients/order"
	productClient "pan/clients/product"
	"pan/dto"
	"pan/model"
	e "pan/utils/errors"

	log "github.com/sirupsen/logrus"
)

type orderService struct{}

type orderServiceInterface interface {
	OrderHistoy(int) (dto.OrdersHistoryDto, e.ApiError)
	NewOrder(NewOrderDto dto.NewOrderDto) (dto.OrderDto, e.ApiError)
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

	var orderAux dto.OrderDto
	var detalleAUX dto.DetalleDto
	var productAUX dto.ProductDto
	var product model.Product
	var order dto.OrdersHistoryDto
	orders = orderCliente.GetOrders(idUser)

	// seteo de []order.orderhistorydto lugar por lugar
	for j := 0; j < len(orders); j++ {
		detalles = orderCliente.GetDetalles(orders[j].Id_Order)
		for k := 0; k < len(detalles); k++ {
			// seteo del producto dentro del detalleAUX a insertar
			product = productClient.GetProductById(detalles[k].Id_Product)
			productAUX.Id_Product = product.Id_Product
			productAUX.Stock = product.Stock
			productAUX.Category = product.Category
			productAUX.Price = product.Price
			productAUX.Name_product = product.Name_product
			detalleAUX.Product = productAUX

			//setup detalle a insertar
			detalleAUX.Cant_det = detalles[k].Cantidad
			detalleAUX.Price_det = detalles[k].Precio_Moment
			orderAux.Det_order = append(orderAux.Det_order, detalleAUX)

		}
		// setup historial con auxiliares
		order.Order = append(order.Order, orderAux)
	}
	order.Id_User = idUser
	return order, nil

}
func (s *orderService) NewOrder(NewOrderDto dto.NewOrderDto) (dto.OrderDto, e.ApiError) {
	var order model.Order
	var detalles []model.Detalle
	var detalleAUX model.Detalle
	var total float32 = 0

	//falta escribir todos los datos en el producto para mostrar bien la order
	for j := 0; j < len(NewOrderDto.Detalles); j++ {
		detalleAUX.Cantidad = NewOrderDto.Detalles[j].Cant_det
		detalleAUX.Id_Product = NewOrderDto.Detalles[j].Product.Id_Product
		detalleAUX.Precio_Moment = NewOrderDto.Detalles[j].Price_det
		detalles = append(detalles, detalleAUX)
		total = total + NewOrderDto.Detalles[j].Price_det*float32(NewOrderDto.Detalles[j].Cant_det)
	}
	order.Total = total
	order.Id_User = NewOrderDto.Id_User
	order = orderCliente.NewOrder(order, detalles)
	log.Debug("newOrderClient Complete")
	var orderDto dto.OrderDto
	if order.Id_Order == 0 {
		return orderDto, e.NewBadRequestApiError("No stock enought")
	}
	//hay una diferencia entre los dto y los modelos ya que los dto tienen datos de otras tablas de la db

	orderDto.Id_order = order.Id_Order
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
