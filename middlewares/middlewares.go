package middlewares

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetUpMiddlewares(e *echo.Echo) {
	// ロガーとリカバリーミドルウェアを使用
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	// CORSを有効化
	// AllowCredentialsをtrueに設定すると、クライアント側でwithCredentialsをtrueに設定する必要がある
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(allowedOrigins, ","),
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowCredentials,
		},
		// ExposeHeaders: []string{
		// 	echo.HeaderSetCookie,
		// },
		AllowCredentials: true,
	}))
}
