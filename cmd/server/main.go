package main

import (
	"github.com/Taylorgtyler/go-rest-api/cmd/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Registering handlers or routes
	handlers.RegisterRoutes(r)

	r.Run(":8080") // listen and serve on :8080
}
