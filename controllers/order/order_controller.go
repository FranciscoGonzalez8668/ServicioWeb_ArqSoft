package orderController

import (
	"net/http"
	service "pan/services"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetHistory(c *gin.Context) {

	log.Debug("Historial de Id_user: " + c.Param("id_user"))
	id_user, _ := strconv.Atoi(c.Param("id_user"))

	ordersHistoryDto, err := service.OrderService.OrderHistoy(id_user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, ordersHistoryDto)

}
