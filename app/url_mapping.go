package app

import (
	//adressController "mvc-go/controllers/adress"
	//productController "mvc-go/controllers/product"
	//userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls(){

	router.GET("/login",//userController.GetUser()) //identificar el usuario ^ devolver una toquen  

	router.GET("/product/:KeyPro",//productController.GetProductByName()) // Busca productos por Palabra clave
	
	router.POST("/compra/add",//ordenController.NewOrder()) //Debe crear la nueva orden verificar cantidades

	router.PUT("/cart",//CartController.CartMananger()) //Sumar Restar sacar y agregar productos en carrito

	router.GET("/User/OrH",//userContrller.GetHistory()) //Devolver compras por user

	router.Get("/product/:KeyCat",//product.Controller.GetProductByCat())//Traer Productos por Categoria


	log.Info("Finishing Mappings configurations")


}