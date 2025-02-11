package handlers_chat_messages

import (
	"net/http"
	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"

	"github.com/labstack/echo"
)

// FetchChatMessagesInRoom は `chat_messages` テーブルと `users` テーブルを結合して、チャットメッセージ情報とユーザー情報を取得する。
func (h *ChatMessagesHandler) FetchChatMessagesInRoom(c echo.Context) error {
	roomId := c.Param("id")
	chatMessages, err := h.ChatMessagesService.FetchChatMessagesInRoom(roomId)

	if err != nil {
		switch err.Error() {
		case "id is required":
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid id",
			})
		case "invalid id":
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid id",
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error fetching chat_messages",
			})
		}
	}

	return c.JSON(http.StatusOK, chatMessages)
}

// CreateChatMessage は `chat_messages` テーブルにメッセージを作成する。
func (h *ChatMessagesHandler) CreateChatMessage(c echo.Context) error {
	var createChatMessageRequest models.CreateChatMessageRequest

	if err := c.Bind(&createChatMessageRequest); err != nil {
		logger.ErrorLog.Printf("Failed to bind create chat_message request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}
	message := createChatMessageRequest.Message
	roomId := createChatMessageRequest.RoomId
	userId := createChatMessageRequest.UserId

	messageId, err := h.ChatMessagesService.CreateChatMessage(message, roomId, userId)
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

	logger.InfoLog.Printf("Created chat_message successfully")
	return c.JSON(http.StatusOK, map[string]string{
		"messageId": messageId,
	})
}
