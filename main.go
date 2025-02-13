package main

import (
	"net/http"
	"nextjs-echo-chat-back-app/config"
	"nextjs-echo-chat-back-app/middlewares"
	"nextjs-echo-chat-back-app/routes"
	"nextjs-echo-chat-back-app/utils/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo"
)

func SetUp(e *echo.Echo) {
	// シグナルハンドラーの設定
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Supabaseのセットアップ
	if err := middlewares.SetUpSupabase(); err != nil {
		logger.ErrorLog.Fatalf("Failed to set up Supabase: %v", err)
	}
	// Supabaseのテストクエリ
	if err := middlewares.TestQuery(); err != nil {
		logger.ErrorLog.Fatalf("Failed to test Supabase connection: %v", err)
	}

	// ミドルウェアの設定
	middlewares.SetUpMiddlewares(e)
	// ルーティングの設定
	routes.SetUpRouter(e)
	// WebSocketのルーティング設定
	routes.SetUpWebSocketRoutes(e)

	go func() {
		<-quit
		logger.InfoLog.Println("Shutting down server...")

		// Supabaseのコネクションプールをクローズ
		middlewares.ClosePool()

		// Echoサーバーのシャットダウン
		if err := e.Close(); err != nil {
			logger.ErrorLog.Printf("Echo shutdown failed: %v", err)
		}

	}()
}

func SetUpWebSocketServer() {
	wsServer := echo.New()
	routes.SetUpWebSocketRoutes(wsServer)

	wsPort := config.WsPort
	if wsPort == "" {
		wsPort = "8081"
	}
	go func() {
		logger.InfoLog.Println("Starting WebSocket server on port 8081")
		if err := wsServer.Start(":" + wsPort); err != nil && err != http.ErrServerClosed {
			logger.ErrorLog.Fatalf("WebSocket server failed: %v", err)
		}
	}()
}

func main() {
	// Echoの初期化
	e := echo.New()

	// WebSocketサーバーのセットアップ
	SetUpWebSocketServer()
	// セットアップ
	SetUp(e)

	// サーバーの起動
	port := config.Port
	if port == "" {
		port = "8080"
	}
	if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
		logger.ErrorLog.Fatalf("Echo server failed: %v", err)
	}
}
