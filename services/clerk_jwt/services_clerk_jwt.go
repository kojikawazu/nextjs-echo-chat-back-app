package services_clerk_jwt

import (
	"errors"
	utils "nextjs-echo-chat-back-app/utils/clerk"
	"nextjs-echo-chat-back-app/utils/logger"
	"strings"

	"github.com/labstack/echo"
)

// CheckClerkToken は `auth_users` テーブルから `userId` を取得する。
func (s *ClerkJwtServiceImpl) CheckClerkToken(c echo.Context) (string, error) {
	// Authorization ヘッダーから JWT を取得
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		logger.ErrorLog.Printf("No authorization header found")
		return "", errors.New("No authorization header found")
	}

	// Bearer トークンを取得
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenStr == authHeader {
		logger.ErrorLog.Printf("Invalid token format")
		return "", errors.New("Invalid token format")
	}

	// JWT の検証と `userId` の取得
	userId, err := utils.VerifyClerkToken(tokenStr)
	if err != nil {
		logger.ErrorLog.Printf("Invalid token: %v", err)
		return "", errors.New("Invalid token")
	}

	return userId, nil
}
