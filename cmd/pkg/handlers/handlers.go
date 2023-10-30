package handlers

import (
	"github.com/Taylorgtyler/go-rest-api/cmd/pkg/middleware"
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

	// Define your routes here and add authentication middleware for routes you want to password protect
	// Middleware can also be implemented globally
	r.POST("/register", RegisterNewUser)
	r.POST("/login", LoginUser)
	r.POST("/activity", middleware.Authenticate(), CreateActivityHandler)
	r.PUT("/activity/:id", middleware.Authenticate(), UpdateActivity(db))
	r.POST("activities", middleware.Authenticate(), InsertMultipleActivitiesHandler(db))
}
