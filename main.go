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

func SetUp(e *echo.Echo, ws *echo.Echo) {
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
	middlewares.IPBlockMiddleware(e)
	// ルーティングの設定
	routes.SetUpRouter(e, ws)

	go func() {
		<-quit
		logger.InfoLog.Println("Shutting down server...")

		// Supabaseのコネクションプールをクローズ
		middlewares.ClosePool()

		// Echoサーバーのシャットダウン
		if err := e.Close(); err != nil {
			logger.ErrorLog.Printf("Echo shutdown failed: %v", err)
		}
		// WebSocketサーバーのシャットダウン
		if err := ws.Close(); err != nil {
			logger.ErrorLog.Printf("WebSocket shutdown failed: %v", err)
		}
	}()
}

func main() {
	// Echoの初期化
	e := echo.New()
	// WebSocketの初期化
	ws := echo.New()

	// セットアップ
	SetUp(e, ws)

	// WebSocketサーバーの起動
	wsPort := config.WsPort
	if wsPort == "" {
		wsPort = "8081"
	}
	go func() {
		if err := ws.Start(":" + wsPort); err != nil && err != http.ErrServerClosed {
			logger.ErrorLog.Fatalf("WebSocket server failed: %v", err)
		}
	}()
	// サーバーの起動
	port := config.Port
	if port == "" {
		port = "8080"
	}
	if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
		logger.ErrorLog.Fatalf("Echo server failed: %v", err)
	}
}
