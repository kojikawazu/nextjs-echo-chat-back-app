package websocket

import (
	"encoding/json"
	"net/http"
	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocketアップグレーダー
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// クライアントの管理
var (
	clients     = make(map[string]map[*websocket.Conn]bool) // roomId → connection map
	clientsLock = sync.Mutex{}
)

// WebSocketの接続を処理
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	logger.InfoLog.Println("WebSocket connection request received")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.ErrorLog.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	logger.InfoLog.Println("Client connected")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logger.ErrorLog.Println("Read Error:", err)
			break
		}

		logger.InfoLog.Println("Received message:", string(msg))

		// クライアントに送信
		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			logger.ErrorLog.Println("Write Error:", err)
			break
		}
	}
}

// WebSocket メッセージをブロードキャスト
func BroadcastMessage(message models.ChatMessages) {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	connections, exists := clients[message.RoomID]
	if !exists {
		return
	}

	msgBytes, err := json.Marshal(message)
	if err != nil {
		logger.ErrorLog.Println("JSON Marshal Error:", err)
		return
	}

	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
			logger.ErrorLog.Println("Write Error:", err)
			conn.Close()
			delete(connections, conn)
		}
	}
}
