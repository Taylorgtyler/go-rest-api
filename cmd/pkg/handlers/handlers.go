package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes defines the routes for the application
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// setting the DB in the context so it can be accessed in the handlers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Define your routes here
	r.POST("/register", RegisterNewUser)
	r.POST("/login", LoginUser)
	r.POST("/activity", CreateActivityHandler)
	r.PUT("/activity/:id", UpdateActivity(db))
	r.POST("activities", InsertMultipleActivitiesHandler(db))
}
