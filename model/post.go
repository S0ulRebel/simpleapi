package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Title  string `gorm:"size:100;not null"`
	Body   string `gorm:"not null"`
}
