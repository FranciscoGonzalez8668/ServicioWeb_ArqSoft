package productController

import (
	"fmt"
	"net/http"
	"pan/dto"
	service "pan/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProductByName(c *gin.Context) {

	log.Debug("Product name to load: " + c.Param("KeyPro"))

	name := c.Param("name")

	var productDto dto.ProductDto

	productDto, err := service.ProductService.GetProductByName(name)

	if err != nil {

		c.JSON(err.Status(), err)
		return

	}

	c.JSON(http.StatusOK, productDto)

}

func GetProductByCat(c *gin.Context) {

	log.Debug("Product category to load: ", c.Param("KeyCat"))
	cat := c.Param("KeyCat")

	productsDto, err := service.ProductService.GetProductByCat(cat)
	fmt.Println("todas las consultas ok")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}
