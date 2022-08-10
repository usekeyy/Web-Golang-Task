package models

import (
	"time"
)

type Task struct {
	ID          int
	Task_detail string `gorm:"type:varchar(100)"`
	Assigne     string `gorm:"type:varchar(100)"`
	Deadline    string
	IsDone      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
