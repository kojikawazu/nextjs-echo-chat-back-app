package models

import "time"

// WebSocket メッセージの構造体
type WebSocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// WebSocket チャットメッセージの構造体
type WebSocketChatMessage struct {
	UserID     string    `json:"user_id"`
	MessageID  string    `json:"message_id"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	LikeCount  int       `json:"like_count"`
	LikedUsers []string  `json:"liked_users"`
}

// WebSocket 参加メッセージの構造体
type WebSocketJoinMessage struct {
	RoomID string `json:"room_id"`
}

// WebSocket エラーメッセージの構造体
type WebSocketErrorMessage struct {
	Error string `json:"error"`
}
