package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	CreatedAt time.Time
}
