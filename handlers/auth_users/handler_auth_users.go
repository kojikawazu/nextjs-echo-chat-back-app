package handlers_auth_users

import (
	"net/http"
	"nextjs-echo-chat-back-app/utils/logger"

	"github.com/labstack/echo"
)

// FetchAuthUsers は `auth_users` テーブルからすべてのユーザー情報を取得する。
func (h *AuthUsersHandler) FetchAuthUsers(c echo.Context) error {

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

	// `auth_users` テーブルからすべてのユーザー情報を取得
	authUsers, err := h.AuthUsersService.FetchAuthUsers()
	if err != nil {
		logger.ErrorLog.Printf("Error fetching auth_users: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error fetching auth_users",
		})
	}

	// 取得したユーザー情報をJSON形式で返す
	logger.InfoLog.Println("Successfully fetched auth_users.")
	return c.JSON(http.StatusOK, authUsers)
}
