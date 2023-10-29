package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	// Your user registration logic here

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}
