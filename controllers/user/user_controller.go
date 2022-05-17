package userController

import (
	"net/http"

	"servicioweb_arqsoft/dto"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	//"strconv"
)

func Login(c *gin.Context) {
	var userdto dto.UserDto
	err := c.BindJSON(&userdto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//falta seguir con cosas del servis que aun no estan implementadas

}
