package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Tasks     []Task         `json:"tasks" gorm:"foreignKey:UserID"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
}
