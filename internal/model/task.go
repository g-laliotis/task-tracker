package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	UserID    uint      `json:"user_id"` // 🔗 Each task belongs to a user
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
