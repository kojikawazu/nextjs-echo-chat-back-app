package models

import "time"

type ChatLikes struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	MessageID string    `json:"message_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
