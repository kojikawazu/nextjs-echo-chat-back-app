package handlers_chat_messages

import (
	"net/http"

	"github.com/labstack/echo"
)

// FetchChatMessagesInRoom は `chat_messages` テーブルと `users` テーブルを結合して、チャットメッセージ情報とユーザー情報を取得する。
func (h *ChatMessagesHandler) FetchChatMessagesInRoom(c echo.Context) error {
	id := c.Param("id")
	chatMessages, err := h.ChatMessagesService.FetchChatMessagesInRoom(id)

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
