package db

import (

	///IMPORTS COPIADOS DEL TEORICO
	/*
		"mvc-go/model"


		_ "github.com/jinzhu/gorm/dialects/mysql"
		log "github.com/sirupsen/logrus"
	*/
	//userClient "mvc-go/clients/user"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	DBName := "PAN"
	DBUser := "root"
	DBPass := "rootroot"
	DBHost := "127.0.0.1"

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3307)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Println(err)
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	//userClient.Db =db

}

func StartDbEngine() {

	//CREAR LAS TABLAS DEL DB
	/*db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Produc{})
	db.AutoMigrate(&model.Order{})*/

	log.Info("Finishing Migration Database Tables")
}
