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

func InsertMultipleActivities(db *gorm.DB, activities []*models.Activity) error {
	if activities == nil {
		err := errors.New("activities slice is nil")
		log.Printf("Error inserting activities: %v\n", err)
		return err
	}

	for _, activity := range activities {
		if activity == nil {
			err := errors.New("an activity is nil")
			log.Printf("Error inserting activity: %v\n", err)
			return err
		}

		if activity.Name == "" {
			err := errors.New("an activity name is empty")
			log.Printf("Error inserting activity: %v\n", err)
			return err
		}

		result := db.Create(&activity)
		if result.Error != nil {
			log.Printf("Error occurred while inserting activity (ID: %d, Name: %s): %v\n", activity.ID, activity.Name, result.Error)
			return result.Error
		}

		log.Printf("Activity inserted successfully with ID: %d\n", activity.ID)
	}

	return nil
}

func UpdateActivity(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extracting the activity ID from the URL parameter.
		activityID := c.Param("id")

		// Finding the activity in the database.
		var activity models.Activity
		if err := db.First(&activity, activityID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
			return
		}

		// Binding the new activity details from the JSON body.
		var newActivity models.Activity
		if err := c.ShouldBindJSON(&newActivity); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Updating the activity.
		if err := db.Model(&activity).Updates(newActivity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Responding with the updated activity.
		c.JSON(http.StatusOK, activity)
	}
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

func InsertMultipleActivitiesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var activities []*models.Activity

		if err := c.ShouldBindJSON(&activities); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := InsertMultipleActivities(db, activities)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Activities inserted successfully"})
	}
}
