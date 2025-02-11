package models

import "time"

type ChatMessages struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	RoomID    string    `json:"room_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateChatMessageRequest struct {
	Message string `json:"message"`
	RoomId  string `json:"room_id"`
	UserId  string `json:"user_id"`
}
