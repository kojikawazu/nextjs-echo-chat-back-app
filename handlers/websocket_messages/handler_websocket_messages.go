package handlers_websocket_messages

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"

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

// WebSocket の接続処理
func (h *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	logger.InfoLog.Println("WebSocket connection request received")

	// WebSocket 接続
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.ErrorLog.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	logger.InfoLog.Println("Client connected, waiting for roomId...")

	// 初回メッセージで `roomId` を受信
	var msg models.WebSocketMessage
	if err := conn.ReadJSON(&msg); err != nil {
		logger.ErrorLog.Println("Failed to read room ID:", err)
		return
	}

	// `join` メッセージでなければエラー
	if msg.Type != "join" {
		logger.ErrorLog.Println("Invalid message type:", msg.Type)
		return
	}

	roomID := msg.Data.RoomID
	logger.InfoLog.Printf("Client joined room.")

	// クライアントを `roomId` に登録
	h.lock.Lock()
	if _, exists := h.clients[roomID]; !exists {
		h.clients[roomID] = make(map[*websocket.Conn]bool)
	}
	h.clients[roomID][conn] = true
	h.lock.Unlock()

	logger.InfoLog.Printf("Client joined room successfully.")

	// メッセージを受信
	for {
		var receivedMessage models.WebSocketMessage
		if err := conn.ReadJSON(&receivedMessage); err != nil {
			logger.ErrorLog.Println("Read Error:", err)
			break
		}

		// `type: message` のときのみ処理
		if receivedMessage.Type == "message" {
			receivedMessage.Data.CreatedAt = time.Now()
			receivedMessage.Data.UpdatedAt = time.Now()

			// メッセージを `roomId` のクライアントに送信
			h.BroadcastMessage(roomID, receivedMessage.Data)
		}
	}

	// クライアント切断時に `roomId` から削除
	h.lock.Lock()
	delete(h.clients[roomID], conn)
	if len(h.clients[roomID]) == 0 {
		delete(h.clients, roomID)
	}
	h.lock.Unlock()
	logger.InfoLog.Println("Client disconnected from room successfully.")
}

// WebSocket メッセージをルームにブロードキャスト
func (h *WebSocketHandler) BroadcastMessage(roomID string, chatMessage models.ChatMessages) {
	logger.InfoLog.Println("Broadcasting message to room.")

	message := models.WebSocketMessage{
		Type: "message",
		Data: chatMessage,
	}

	h.lock.Lock()
	defer h.lock.Unlock()

	// ルームIDに紐づくクライアントを取得
	connections, exists := h.clients[roomID]
	if !exists {
		logger.ErrorLog.Println("Room not found.")
		return
	}

	// メッセージを JSON に変換
	msgBytes, err := json.Marshal(message)
	if err != nil {
		logger.ErrorLog.Println("JSON Marshal Error:", err)
		return
	}

	// 各クライアントに送信
	logger.InfoLog.Println("Sending message to clients.")
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
			logger.ErrorLog.Println("Write Error:", err)
			conn.Close()
			delete(connections, conn)
		}
	}

	logger.InfoLog.Println("Message broadcasted successfully.")
}
