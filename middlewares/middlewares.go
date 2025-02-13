package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ブロックする IP アドレスのリスト
var blockedIPs = make(map[string]bool)

func init() {
	ips := os.Getenv("BLOCKED_IP_ADDRESSES")
	if ips != "" {
		for _, ip := range strings.Split(ips, ",") {
			blockedIPs[strings.TrimSpace(ip)] = true
		}
	}
}

// IPアドレスをブロックするミドルウェア
func IPBlockMiddleware(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			clientIP := c.RealIP()
			if blockedIPs[clientIP] {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "Access denied",
				})
			}
			return next(c)
		}
	})
}

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
