package productController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"pan/dto"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProductByName(c *gin.Context) {

	log.Debug("Product name to load: " + c.Param("name"))

	name, _ := strconv.Atoi(c.Param("name"))

	var productDto dto.ProductDto

	productDto, err := service.ProductService.GetProductByName(name)

	if err != nil {

		c.JSON(err.Status(), err)
		return

	}

	c.JSON(http.StatusOK, productDto)

}

func GetProductByCat(c *gin.Context) {

	c.JSON(http.StatusOK, "")
}
