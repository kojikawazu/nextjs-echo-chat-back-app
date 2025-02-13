package models

// WebSocket メッセージの構造体
type WebSocketMessage struct {
	Type string       `json:"type"`
	Data ChatMessages `json:"data"`
}
