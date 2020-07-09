package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jcabrera/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)
func StartApplication(){
	mapUrls()
	logger.Info("Starting app")
	router.Run(":8080")
}