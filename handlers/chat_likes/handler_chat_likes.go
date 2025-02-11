package handlers_chat_likes

import (
	"net/http"

	"github.com/labstack/echo"
)

// FetchChatLikesInUsers は `chat_likes` テーブルと `users` テーブルを結合して、チャットいいね情報とユーザー情報を取得する。
func (h *ChatLikesHandler) FetchChatLikesInUsers(c echo.Context) error {
	id := c.Param("id")
	chatLikes, err := h.ChatLikesService.FetchChatLikesInUsers(id)

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
				"error": "Error fetching chat_likes",
			})
		}
	}

	return c.JSON(http.StatusOK, chatLikes)
}
