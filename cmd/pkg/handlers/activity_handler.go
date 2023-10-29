package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/Taylorgtyler/go-rest-api/cmd/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InsertActivity(db *gorm.DB, activity *models.Activity) error {
	if activity == nil {
		err := errors.New("activity is nil")
		log.Printf("Error inserting activity: %v\n", err)
		return err
	}

	if activity.Name == "" {
		err := errors.New("activity name is empty")
		log.Printf("Error inserting activity: %v\n", err)
		return err
	}

	result := db.Create(&activity)
	if result.Error != nil {
		log.Printf("Error occurred while inserting activity: %v\n", result.Error)
		return result.Error
	}

	log.Printf("Activity inserted successfully with ID: %d\n", activity.ID)
	return nil
}

func UpdateActivity(r *gin.Engine) {

}

func CreateActivityHandler(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB) // Get db instance from context

	var activity models.Activity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := InsertActivity(db, &activity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity inserted successfully", "activity": activity})
}
