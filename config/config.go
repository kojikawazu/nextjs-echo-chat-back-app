package config

import (
	"nextjs-echo-chat-back-app/utils/logger"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	JwtKey       []byte
	IsProduction bool
	Port         string
	WsPort       string
	once         sync.Once
)

func init() {
	once.Do(func() {
		// ログ設定の初期化
		logger.SetUpLogger()

		// 環境変数の読み込み
		err := godotenv.Load()
		if err != nil {
			logger.ErrorLog.Printf("error loading .env file: %v", err)
			logger.ErrorLog.Fatal("Failed to load .env file")
		}

		// 環境変数をグローバル変数にセット
		JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
		IsProduction = os.Getenv("ENV") == "production"
		WsPort = os.Getenv("WS_PORT")
		Port = os.Getenv("PORT")

		// 環境変数の読み込み結果をログに出力（デバッグ用）
		logger.InfoLog.Printf("Config Loaded: IsProduction=%v\n", IsProduction)
	})
}
