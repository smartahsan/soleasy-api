// main.go
package main

import (
	"soleasy/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
