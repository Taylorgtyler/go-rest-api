package models

import (
	"time"

	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	ID        uint
	Name      string
	UserID    int
	User      User
	CreatedAt time.Time
}
