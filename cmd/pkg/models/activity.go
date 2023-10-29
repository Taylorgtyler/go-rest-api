package models

import (
	"database/sql"
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
	ListID      sql.NullInt64 `gorm:"default:NULL"`
	List        List
}
