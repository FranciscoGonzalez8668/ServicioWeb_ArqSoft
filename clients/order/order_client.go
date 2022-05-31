package product

import (
	"fmt"

	"pan/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHistory(id_user int) ([]model.OrderDet, []int) {
	var orderHistory []model.OrderDet

	var orderAux model.OrderDet
	for j := 0; j < 4; j++ {
		fmt.Println("se entra ford client i= ", j)
		orderAux.Id_order = j
		orderAux.Total = 5
		orderAux.Cantidad = j * 2
		orderAux.Precio_Moment = 100
		orderAux.Product = "Led"
		/*for i := 0; i < 2; i++ {
			orderAux.Detalles[i].Cantidad = 2
			orderAux.Detalles[i].Id_detalle = i
			orderAux.Detalles[i].Product_Name = "jabon"
			orderAux.Detalles[i].Precio_Moment = 2
		}*/
		orderHistory = append(orderHistory, orderAux)
	}

	//var detalles []model.Detalle

	//SQL necesario
	//SELECT s.staff_id, p.amount, p.payment_id ,p.last_update
	//FROM staff s INNER JOIN payment p on p.staff_id = s.staff_id;
	/*
		Db.Model(&model.Order{}).Select("order.Id_Order,detalle.Cantidad, detalle.Precio_Moment, product.Name_product").
		Joins("left join detalle on detalle.Id_Order = order.Id_Order").
		Joins("left join product on product.Id_Product = detalle.Id_Product").Find(&orderHistory)
	*/

	var id_order []int
	Db.Model(&model.Order{}).Select("order.Id_Order").Joins("left join user on user.Id_Order = order.Id_order").Group("order.Id_Order").Find(&id_order)
	return orderHistory, id_order

}

func NewOrder(order model.Order) model.Order {
	result := Db.Create(&order)

	if result.Error != nil {
		log.Error("no se pudo crear en la base de datos")
	}
	log.Debug("order created: ", order.Id_order)
	return order
}
