package routes

import (
	"nextjs-echo-chat-back-app/websocket"

	"github.com/labstack/echo"
)

// WebSocketのルーティング設定
func SetUpWebSocketRoutes(e *echo.Echo) {
	e.GET("/ws", func(c echo.Context) error {
		websocket.HandleWebSocket(c.Response(), c.Request())
		return nil
	})
}
