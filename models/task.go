package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	ID          uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title       string     `json:"title" gorm:"not null"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status" gorm:"default:'todo'"`
	DueDate     *time.Time `json:"due_date"`
	Priority    int        `json:"priority" gorm:"default:1"`
	UserID      uuid.UUID  `json:"user_id" gorm:"type:uuid;not null"`
	User        User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (task *Task) BeforeCreate(tx *gorm.DB) error {
	if task.ID == uuid.Nil {
		task.ID = uuid.New()
	}
	return nil
}
