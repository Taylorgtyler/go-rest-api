package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID            uint
	UserID        int
	User          User
	ScheduledTime time.Time
}
