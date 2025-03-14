package handlers_chat_likes

import (
	"net/http"
	"nextjs-echo-chat-back-app/utils/logger"

	"github.com/labstack/echo"
)

// FetchChatLikesInUsers は `chat_likes` テーブルと `users` テーブルを結合して、チャットいいね情報とユーザー情報を取得する。
func (h *ChatLikesHandler) FetchChatLikesInUsers(c echo.Context) error {

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

	encryptedMessageId := c.Param("id")
	chatLikes, err := h.ChatLikesService.FetchChatLikesInUsers(encryptedMessageId)

	if err != nil {
		switch err.Error() {
		case "messageId is required":
			logger.ErrorLog.Printf("Invalid messageId: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid messageId",
			})
		case "invalid messageId":
			logger.ErrorLog.Printf("Invalid messageId: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid messageId",
			})
		default:
			logger.ErrorLog.Printf("Error fetching chat_likes: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error fetching chat_likes",
			})
		}
	}

	logger.InfoLog.Printf("Fetched chat_likes successfully")
	return c.JSON(http.StatusOK, chatLikes)
}

// CreateChatLike は `chat_likes` テーブルに新しいいいねを作成する。
func (h *ChatLikesHandler) CreateChatLike(c echo.Context) error {
	encryptedMessageId := c.Param("id")

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

	// いいねの作成
	likeId, err := h.ChatLikesService.CreateChatLike(encryptedMessageId, userId)
	if err != nil {
		switch err.Error() {
		case "messageId is required":
			logger.ErrorLog.Printf("Invalid messageId: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid messageId",
			})
		case "userId is required":
			logger.ErrorLog.Printf("Invalid userId: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid userId",
			})
		case "invalid messageId":
			logger.ErrorLog.Printf("Invalid messageId: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid messageId",
			})
		default:
			logger.ErrorLog.Printf("Error creating chat_like: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error creating chat_like",
			})
		}
	}

	response := map[string]string{
		"id": likeId,
	}

	logger.InfoLog.Printf("Created chat_like successfully")
	return c.JSON(http.StatusOK, response)
}

// DeleteChatLike は `chat_likes` テーブルからいいねを削除する。
func (h *ChatLikesHandler) DeleteChatLike(c echo.Context) error {
	encryptedMessageId := c.Param("id")

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

	// いいねの削除
	likeId, err := h.ChatLikesService.DeleteChatLike(encryptedMessageId, userId)
	if err != nil {
		switch err.Error() {
		case "messageId is required":
			logger.ErrorLog.Printf("Invalid messageId: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid messageId",
			})
		case "userId is required":
			logger.ErrorLog.Printf("Invalid userId: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid userId",
			})
		case "invalid messageId":
			logger.ErrorLog.Printf("Invalid messageId: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid messageId",
			})
		default:
			logger.ErrorLog.Printf("Error deleting chat_like: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error deleting chat_like",
			})
		}
	}

	response := map[string]string{
		"id": likeId,
	}

	logger.InfoLog.Printf("Deleted chat_like successfully")
	return c.JSON(http.StatusOK, response)
}
