package handlers_websocket_messages

import (
	"encoding/json"
	"net/http"
	"time"

	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"

	"github.com/gorilla/websocket"
)

// WebSocket の接続処理
func (h *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	logger.InfoLog.Println("WebSocket connection request received")

	// WebSocket 接続
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.ErrorLog.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer func() {
		conn.Close()
		h.removeClient(conn)
	}()

	logger.InfoLog.Println("Client connected, waiting for roomId...")

	// 初回メッセージで `roomId` を受信
	var msg models.WebSocketMessage
	if err := conn.ReadJSON(&msg); err != nil {
		h.sendErrorAndClose(conn, "Failed to read initial message")
		logger.ErrorLog.Printf("Failed to read initial message: %v", err)
		return
	}
	if msg.Type != "join" {
		h.sendErrorAndClose(conn, "First message must be join type")
		logger.ErrorLog.Printf("First message not join type: %v", msg.Type)
		return
	}

	logger.InfoLog.Println("Received join message.")

	// Dataを型安全にキャスト
	joinDataMap, ok := msg.Data.(map[string]interface{})
	if !ok {
		h.sendErrorAndClose(conn, "Invalid join data format")
		logger.ErrorLog.Printf("Invalid join data format: %v", msg.Data)
		return
	}
	roomIDRaw, exists := joinDataMap["room_id"]
	if !exists {
		h.sendErrorAndClose(conn, "room_id key missing in join data")
		logger.ErrorLog.Printf("room_id key missing: %+v", joinDataMap)
		return
	}
	roomID, ok := roomIDRaw.(string)
	if !ok || roomID == "" {
		h.sendErrorAndClose(conn, "room_id must be a non-empty string")
		logger.ErrorLog.Printf("Invalid room_id value: %+v", roomIDRaw)
		return
	}

	// クライアントを追加
	h.addClient(roomID, conn)
	logger.InfoLog.Println("Client joined room.")

	for {
		// メッセージを受信
		var receivedMessage models.WebSocketMessage
		if err := conn.ReadJSON(&receivedMessage); err != nil {
			logger.ErrorLog.Printf("Read Error: %v", err)
			break
		}

		logger.InfoLog.Println("Received message.")

		// メッセージの型に応じて処理を分岐
		switch receivedMessage.Type {
		case "message":
			logger.InfoLog.Println("Received message.")

			messageDataMap, ok := receivedMessage.Data.(map[string]interface{})
			if !ok {
				h.sendError(conn, "Invalid message format")
				logger.ErrorLog.Printf("Invalid message format: %v", receivedMessage.Data)
				continue
			}

			message, ok := messageDataMap["message"].(string)
			if !ok || message == "" {
				h.sendError(conn, "Message content cannot be empty")
				logger.ErrorLog.Printf("Message content cannot be empty: %v", messageDataMap)
				continue
			}

			chatMessage := models.WebSocketChatMessage{
				UserID:    messageDataMap["user_id"].(string),
				MessageID: messageDataMap["message_id"].(string),
				Name:      messageDataMap["name"].(string),
				Content:   message,
				CreatedAt: time.Now(),
			}

			// メッセージをブロードキャスト
			h.BroadcastMessage(roomID, chatMessage)

		default:
			h.sendError(conn, "Unsupported message type")
			logger.ErrorLog.Printf("Unsupported message type: %v", receivedMessage.Type)
		}
	}

	logger.InfoLog.Println("Client disconnected from room successfully.")
}

// エラー送信（接続を切断しない）
func (h *WebSocketHandler) sendError(conn *websocket.Conn, errMsg string) {
	errMessage := models.WebSocketMessage{
		Type: "error",
		Data: models.WebSocketErrorMessage{Error: errMsg},
	}
	if err := conn.WriteJSON(errMessage); err != nil {
		logger.ErrorLog.Printf("Failed to send error message: %v", err)
	}
}

// エラー送信後、接続をクローズする
func (h *WebSocketHandler) sendErrorAndClose(conn *websocket.Conn, errMsg string) {
	h.sendError(conn, errMsg)
	conn.Close()
	h.removeClient(conn)
}

// クライアントを安全に追加する
func (h *WebSocketHandler) addClient(roomID string, conn *websocket.Conn) {
	h.lock.Lock()
	defer h.lock.Unlock()

	if _, exists := h.clients[roomID]; !exists {
		h.clients[roomID] = make(map[*websocket.Conn]bool)
	}
	h.clients[roomID][conn] = true
}

// クライアントを安全に削除する
func (h *WebSocketHandler) removeClient(conn *websocket.Conn) {
	h.lock.Lock()
	defer h.lock.Unlock()

	for roomID, conns := range h.clients {
		if _, exists := conns[conn]; exists {
			delete(conns, conn)
			logger.InfoLog.Printf("Removed client from room: %s", roomID)
			if len(conns) == 0 {
				delete(h.clients, roomID)
				logger.InfoLog.Printf("Deleted empty room: %s", roomID)
			}
			break
		}
	}
}

// WebSocket メッセージをルームにブロードキャスト
func (h *WebSocketHandler) BroadcastMessage(roomID string, chatMessage models.WebSocketChatMessage) {
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
