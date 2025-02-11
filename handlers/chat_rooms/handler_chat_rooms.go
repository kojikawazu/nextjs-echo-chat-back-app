package handlers_chat_rooms

import (
	"net/http"

	"github.com/labstack/echo"
)

// FetchChatRooms は `chat_rooms` テーブルからすべてのチャットルーム情報を取得する。
func (h *ChatRoomsHandler) FetchChatRooms(c echo.Context) error {
	chatRooms, err := h.ChatRoomsService.FetchChatRooms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error fetching chat_rooms",
		})
	}
	return c.JSON(http.StatusOK, chatRooms)
}

// FetchUsersInRoom は `chat_rooms` テーブルと `users` テーブルを結合して、チャットルーム情報とユーザー情報を取得する。
func (h *ChatRoomsHandler) FetchUsersInRoom(c echo.Context) error {
	id := c.Param("id")
	chatRooms, err := h.ChatRoomsService.FetchUsersInRoom(id)

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
				"error": "Error fetching chat_rooms with users",
			})
		}
	}

	return c.JSON(http.StatusOK, chatRooms)
}
