package models

import "time"

// WebSocket メッセージの構造体
type WebSocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// WebSocket チャットメッセージの構造体
type WebSocketChatMessage struct {
	RoomID    string    `json:"roomId"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// WebSocket 参加メッセージの構造体
type WebSocketJoinMessage struct {
	RoomID string `json:"roomId"`
}

// WebSocket エラーメッセージの構造体
type WebSocketErrorMessage struct {
	Error string `json:"error"`
}
