package handlers_chat_rooms

import (
	"net/http"
	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"

	"github.com/labstack/echo"
)

// FetchChatRooms は `chat_rooms` テーブルからすべてのチャットルーム情報を取得する。
func (h *ChatRoomsHandler) FetchChatRooms(c echo.Context) error {

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

	chatRooms, err := h.ChatRoomsService.FetchChatRooms()
	if err != nil {
		logger.ErrorLog.Printf("Failed to fetch chat_rooms: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error fetching chat_rooms",
		})
	}

	logger.InfoLog.Println("Fetched chat_rooms.")
	return c.JSON(http.StatusOK, chatRooms)
}

// FetchUsersInRoom は `chat_rooms` テーブルと `users` テーブルを結合して、チャットルーム情報とユーザー情報を取得する。
func (h *ChatRoomsHandler) FetchUsersInRoom(c echo.Context) error {

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

	id := c.Param("id")
	chatRooms, err := h.ChatRoomsService.FetchUsersInRoom(id)

	if err != nil {
		switch err.Error() {
		case "id is required":
			logger.ErrorLog.Printf("Failed to fetch users in room: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid id",
			})
		case "invalid id":
			logger.ErrorLog.Printf("Failed to fetch users in room: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid id",
			})
		default:
			logger.ErrorLog.Printf("Failed to fetch users in room: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error fetching chat_rooms with users",
			})
		}
	}

	logger.InfoLog.Println("Fetched users in room.")
	return c.JSON(http.StatusOK, chatRooms)
}

// CreateRoom は新しいチャットルームを作成する。
func (h *ChatRoomsHandler) CreateRoom(c echo.Context) error {
	var createRoomRequest models.CreateRoomRequest

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

	if err := c.Bind(&createRoomRequest); err != nil {
		logger.ErrorLog.Printf("Failed to bind create room request: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}
	roomName := createRoomRequest.RoomName

	roomId, err := h.ChatRoomsService.CreateRoom(roomName)
	if err != nil {
		switch err.Error() {
		case "room name is required":
			logger.ErrorLog.Printf("Failed to create room: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid room name",
			})
		default:
			logger.ErrorLog.Printf("Failed to create room: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error creating room",
			})
		}
	}

	logger.InfoLog.Println("Created room.")
	return c.JSON(http.StatusOK, map[string]string{
		"room_id": roomId,
	})
}
