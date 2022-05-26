package userController

import (
	//"fmt"
	"net/http"
	"pan/dto"

	service "pan/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	//"strconv"
)

var jwtKey = []byte("secret_key")

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func LoginUser(c *gin.Context) {
	var loginDto dto.LoginDto
	err := c.BindJSON(&loginDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	tokenDto, er := service.UserService.LoginUser(loginDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	tkn, err := jwt.Parse(tokenDto.Token, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, "Invalid Signature")
			return
		}
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, "Invalid Token")
		return
	}
	c.JSON(http.StatusCreated, "Login Succesful")
	//falta seguir con cosas del servis que aun no estan implementadas

}
