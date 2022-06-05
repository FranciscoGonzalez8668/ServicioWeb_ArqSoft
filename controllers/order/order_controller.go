package orderController

import (
	"net/http"
	"pan/dto"
	service "pan/services"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetHistory(c *gin.Context) {

	id_user, _ := strconv.Atoi(c.Param("IU"))

	ordersHistoryDto, err := service.OrderService.OrderHistoy(id_user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, ordersHistoryDto)

}

func NewOrder(c *gin.Context) {
	var NewOrderDto dto.NewOrderDto
	err := c.BindJSON(&NewOrderDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	order, er := service.OrderService.NewOrder(NewOrderDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, order)

}
