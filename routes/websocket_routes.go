package routes

import (
	handlers_websocket_messages "nextjs-echo-chat-back-app/handlers/websocket_messages"

	"github.com/labstack/echo"
)

// WebSocketのルーティング設定
func SetUpWebSocketRoutes(e *echo.Echo) {

	// WebSocket ハンドラの作成
	websocketHandler := handlers_websocket_messages.NewWebSocketHandler()

	// WebSocket 接続
	e.GET("/ws", func(c echo.Context) error {
		websocketHandler.HandleWebSocket(c.Response(), c.Request())
		return nil
	})
}
