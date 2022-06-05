package productController

import (
	"fmt"
	"net/http"
	"pan/dto"
	service "pan/services"

	"github.com/gin-gonic/gin"
)

func GetProductByName(c *gin.Context) {

	name := c.Param("KeyPro")

	var productDto dto.ProductsDto

	productDto, err := service.ProductService.GetProductByName(name)

	if err != nil {

		c.JSON(err.Status(), err)
		return

	}

	c.JSON(http.StatusOK, productDto)

}

func GetProductByCat(c *gin.Context) {

	cat := c.Param("KeyCat")

	productsDto, err := service.ProductService.GetProductByCat(cat)
	fmt.Println("todas las consultas ok")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}
