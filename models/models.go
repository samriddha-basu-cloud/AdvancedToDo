package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Priority    int       `json:"priority"` // 0 - Due today, 1 - Due in 1-2 days, 2 - Due in 3-4 days, 3 - Due in 5+ days
	Status      string    `json:"status"`   // "TODO", "IN_PROGRESS", "DONE"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type SubTask struct {
	ID        int       `json:"id"`
	TaskID    int       `json:"task_id"`
	Status    int       `json:"status"` // 0 - Incomplete, 1 - Complete
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type User struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Priority    int    `json:"priority"` // 0, 1, 2 for Twilio calling priority
}
