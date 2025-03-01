package handlers_websocket_messages

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocketHandler 構造体
type WebSocketHandler struct {
	upgrader websocket.Upgrader
	clients  map[string]map[*websocket.Conn]bool // roomId -> clients
	lock     sync.Mutex
}

// NewWebSocketHandler: DI で WebSocketHandler を作成
func NewWebSocketHandler() *WebSocketHandler {
	return &WebSocketHandler{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		clients: make(map[string]map[*websocket.Conn]bool),
	}
}
