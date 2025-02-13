package services_clerk_jwt

import "github.com/labstack/echo"

// ClerkJwtServiceインターフェース
type ClerkJwtService interface {
	CheckClerkToken(c echo.Context) (string, error)
}

// ClerkJwtServiceImpl は ClerkJwtService の実装
type ClerkJwtServiceImpl struct {
}

// NewClerkJwtService は ClerkJwtService の新しいインスタンスを作成する。
func NewClerkJwtService() ClerkJwtService {
	return &ClerkJwtServiceImpl{}
}
