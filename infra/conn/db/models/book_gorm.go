package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID              uint
	Title           string
	Author          string
	PublicationYear string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
