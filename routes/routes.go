package routes

import (
	"soleasy/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// getTransaction route
	router.GET("/getTransaction/:signature", handlers.GetTransaction)

	router.GET("/getBlockNumber", handlers.GetBlockNumber)

}
