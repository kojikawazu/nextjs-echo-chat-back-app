package handlers_chat_messages

import (
	"net/http"
	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"

	"github.com/labstack/echo"
)

// FetchChatMessagesInRoom は `chat_messages` テーブルと `users` テーブルを結合して、チャットメッセージ情報とユーザー情報を取得する。
func (h *ChatMessagesHandler) FetchChatMessagesInRoom(c echo.Context) error {

	// Authorization ヘッダーから JWT を取得
	_, err := h.ClerkJwtService.CheckClerkToken(c)
	if err != nil {
		switch err.Error() {
		case "No authorization header found":
			logger.ErrorLog.Printf("No authorization header found")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		case "Invalid token format":
			logger.ErrorLog.Printf("Invalid token format")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
		case "Invalid token":
			logger.ErrorLog.Printf("Invalid token")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}
	}

	encryptedRoomId := c.Param("id")
	chatMessages, err := h.ChatMessagesService.FetchChatMessagesInRoom(encryptedRoomId)

	if err != nil {
		switch err.Error() {
		case "id is required":
			logger.ErrorLog.Printf("Invalid id: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid id",
			})
		case "invalid id":
			logger.ErrorLog.Printf("Invalid id: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid id",
			})
		default:
			logger.ErrorLog.Printf("Error fetching chat_messages: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error fetching chat_messages",
			})
		}
	}

	logger.InfoLog.Println("Successfully fetched chat_messages.")
	return c.JSON(http.StatusOK, chatMessages)
}

// CreateChatMessage は `chat_messages` テーブルにメッセージを作成する。
func (h *ChatMessagesHandler) CreateChatMessage(c echo.Context) error {
	var createChatMessageRequest models.CreateChatMessageRequest

	// Authorization ヘッダーから JWT を取得
	userId, err := h.ClerkJwtService.CheckClerkToken(c)
	if err != nil {
		switch err.Error() {
		case "No authorization header found":
			logger.ErrorLog.Printf("No authorization header found")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		case "Invalid token format":
			logger.ErrorLog.Printf("Invalid token format")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
		case "Invalid token":
			logger.ErrorLog.Printf("Invalid token")
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}
	}

	// リクエストボディのバインド
	if err := c.Bind(&createChatMessageRequest); err != nil {
		logger.ErrorLog.Printf("Failed to bind create chat_message request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	message := createChatMessageRequest.Message
	roomId := createChatMessageRequest.RoomId

	// チャットメッセージの作成
	msg, err := h.ChatMessagesService.CreateChatMessage(message, roomId, userId)
	if err != nil {
		switch err.Error() {
		case "message, roomId, userId is required":
			logger.ErrorLog.Printf("Failed to create chat_message: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request",
			})
		case "invalid id":
			logger.ErrorLog.Printf("Failed to create chat_message: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid id",
			})
		default:
			logger.ErrorLog.Printf("Failed to create chat_message: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error creating chat_message",
			})
		}
	}

	// WebSocketでメッセージを送信
	h.WebSocketHandler.BroadcastMessage(roomId, msg)

	logger.InfoLog.Printf("Created chat_message successfully")
	return c.JSON(http.StatusOK, map[string]string{
		"messageId": msg.ID,
	})
}
