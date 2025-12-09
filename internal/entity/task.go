package entity

import "time"

type Task struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"password"`
	UserId    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
