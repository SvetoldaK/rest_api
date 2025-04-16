package tasksService

import (
	"awesomeProject/internal/userService"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Task   string           `json:"task"`
	IsDone bool             `json:"is_done"`
	UserID uint             `json:"user_id" gorm:"not null"`
	User   userService.User `json:"user" gorm:"foreignKey:UserID"`
}
