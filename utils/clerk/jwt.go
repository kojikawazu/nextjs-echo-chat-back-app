package utils

import (
	"fmt"
	"time"

	"nextjs-echo-chat-back-app/utils/logger"

	"github.com/golang-jwt/jwt/v4"
)

// JWTトークンを検証し、userIdを取得
func VerifyClerkToken(tokenStr string) (string, error) {
	// JWTのパース
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// `kid` の存在チェック
		kid, ok := token.Header["kid"].(string)
		if !ok {
			logger.ErrorLog.Printf("missing kid in token header")
			return nil, fmt.Errorf("missing kid in token header")
		}

		// `kid` に基づいて Clerk の公開鍵を取得
		publicKey, err := GetClerkPublicKey(kid)
		if err != nil {
			logger.ErrorLog.Printf("failed to get Clerk public key: %v", err)
			return nil, fmt.Errorf("failed to get Clerk public key: %v", err)
		}
		return publicKey, nil
	})

	if err != nil {
		logger.ErrorLog.Printf("Failed to parse JWT: %v", err)
		return "", fmt.Errorf("invalid token: %v", err)
	}

	// クレーム（Claims）を取得
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// `sub`（ユーザーID）を取得
		userId, ok := claims["sub"].(string)
		if !ok {
			logger.ErrorLog.Printf("invalid token: missing userId")
			return "", fmt.Errorf("invalid token: missing userId")
		}

		// トークンの有効期限チェック
		exp, ok := claims["exp"].(float64)
		if !ok {
			logger.ErrorLog.Printf("invalid token: missing expiration")
			return "", fmt.Errorf("invalid token: missing expiration")
		}
		if time.Now().Unix() > int64(exp) {
			logger.ErrorLog.Printf("token expired")
			return "", fmt.Errorf("token expired")
		}

		return userId, nil
	}

	logger.ErrorLog.Printf("invalid token")
	return "", fmt.Errorf("invalid token")
}
