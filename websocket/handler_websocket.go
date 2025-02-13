package websocket

import (
	"encoding/json"
	"net/http"
	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocket アップグレーダー
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// クライアントの管理 (roomId → connection map)
var (
	clients     = make(map[string]map[*websocket.Conn]bool)
	clientsLock = sync.Mutex{}
)

// WebSocket メッセージの構造体
type WebSocketMessage struct {
	Type string              `json:"type"`
	Data models.ChatMessages `json:"data"`
}

// WebSocket の接続処理
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	logger.InfoLog.Println("WebSocket connection request received")

	// WebSocket 接続
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.ErrorLog.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	logger.InfoLog.Println("Client connected, waiting for roomId...")

	// 初回メッセージで `roomId` を受信
	var msg WebSocketMessage
	if err := conn.ReadJSON(&msg); err != nil {
		logger.ErrorLog.Println("Failed to read room ID:", err)
		return
	}

	// 初回メッセージで `join` を受信
	if msg.Type != "join" {
		logger.ErrorLog.Println("Invalid message type:", msg.Type)
		return
	}

	roomID := msg.Data.RoomID
	userID := msg.Data.UserID
	logger.InfoLog.Printf("Client joined room: %s (User: %s)", roomID, userID)

	// クライアントを `roomId` に登録
	clientsLock.Lock()
	if _, exists := clients[roomID]; !exists {
		clients[roomID] = make(map[*websocket.Conn]bool)
	}
	clients[roomID][conn] = true
	clientsLock.Unlock()

	// メッセージを受信
	for {
		var receivedMessage WebSocketMessage
		if err := conn.ReadJSON(&receivedMessage); err != nil {
			logger.ErrorLog.Println("Read Error:", err)
			break
		}

		// `type: message` のときのみ処理
		if receivedMessage.Type == "message" {
			receivedMessage.Data.CreatedAt = time.Now()
			receivedMessage.Data.UpdatedAt = time.Now()

			// メッセージを `roomId` のクライアントに送信
			BroadcastMessage(roomID, receivedMessage.Data)
		}
	}

	// クライアント切断時に `roomId` から削除
	clientsLock.Lock()
	delete(clients[roomID], conn)
	if len(clients[roomID]) == 0 {
		delete(clients, roomID)
	}
	clientsLock.Unlock()
	logger.InfoLog.Printf("Client disconnected from room: %s (User: %s)", roomID, userID)
}

// WebSocket メッセージをルームにブロードキャスト
func BroadcastMessage(roomID string, chatMessage models.ChatMessages) {
	message := WebSocketMessage{
		Type: "message",
		Data: chatMessage,
	}

	clientsLock.Lock()
	defer clientsLock.Unlock()

	// ルームIDに紐づくクライアントを取得
	connections, exists := clients[roomID]
	if !exists {
		logger.ErrorLog.Println("Room not found:", roomID)
		return
	}

	// メッセージを JSON に変換
	msgBytes, err := json.Marshal(message)
	if err != nil {
		logger.ErrorLog.Println("JSON Marshal Error:", err)
		return
	}

	// 各クライアントに送信
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
			logger.ErrorLog.Println("Write Error:", err)
			conn.Close()
			delete(connections, conn)
		}
	}
}
