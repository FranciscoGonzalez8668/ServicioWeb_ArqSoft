package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartRoute() {
	mapUrls()

	log.Info("Starting server")

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Run(":3307")
}
