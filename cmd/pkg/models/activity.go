package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	ID          uint
	Name        string
	Description string
	CreatedAt   time.Time
	DueDate     time.Time
	Completed   bool
	ListID      int
	List        List
}
