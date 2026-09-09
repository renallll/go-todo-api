package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`

	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"-"`
}
