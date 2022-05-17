package app

import (
	//adressController "mvc-go/controllers/adress"
	productController "servicioweb_arqsoft/controllers/product"
	userController "servicioweb_arqsoft/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	router.GET("/login", userController.GetUser) //identificar el usuario ^ devolver una toquen

	router.GET("/product/product/:KeyPro", productController.GetProductByName) // Busca productos por Palabra clave

	router.POST("/compra/add") //ordenController.NewOrder) //Debe crear la nueva orden verificar cantidades

	router.PUT("/cart") //CartController.CartMananger) //Sumar Restar sacar y agregar productos en carrito

	router.GET("/User/OrH") //userContrller.GetHistory) //Devolver compras por user

	router.GET("/product/cat/:KeyCat") //product.Controller.GetProductByCat)//Traer Productos por Categoria

	log.Info("Finishing mappings configurations")

}
