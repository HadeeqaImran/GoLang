package models

import "time"

type Status string

const (
	TODO  Status = "todo"
	DOING Status = "doing"
	DONE  Status = "done"
)

type Task struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}
