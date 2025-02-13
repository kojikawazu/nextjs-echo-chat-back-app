package handlers_auth_users

import (
	"net/http"
	utils "nextjs-echo-chat-back-app/utils/clerk"
	"nextjs-echo-chat-back-app/utils/logger"
	"strings"

	"github.com/labstack/echo"
)

// FetchAuthUsers は `auth_users` テーブルからすべてのユーザー情報を取得する。
func (h *AuthUsersHandler) FetchAuthUsers(c echo.Context) error {

	// Authorization ヘッダーから JWT を取得
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		logger.ErrorLog.Printf("No authorization header found")
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	// Bearer トークンを取得
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == authHeader {
		logger.ErrorLog.Printf("Invalid token format")
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
	}

	// JWT の検証と `userId` の取得
	_, err := utils.VerifyClerkToken(tokenStr)
	if err != nil {
		logger.ErrorLog.Printf("Invalid token: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}

	authUsers, err := h.AuthUsersService.FetchAuthUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error fetching auth_users",
		})
	}
	return c.JSON(http.StatusOK, authUsers)
}
