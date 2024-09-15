// main.go
package main

import (
	"soleasy/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
