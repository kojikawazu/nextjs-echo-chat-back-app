package middlewares

import (
	"context"
	"fmt"
	"nextjs-echo-chat-back-app/utils/logger"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	// Supabaseとのやり取りに使用するグローバルなコンテキスト。
	Ctx = context.Background()
	// Supabaseとの接続プールです。クエリ実行時に使用。
	Pool *pgxpool.Pool
	// コネクションプールのクローズを1度だけ行うためのsync.Once
	closeOnce sync.Once
)

// Supabaseの接続を初期化
// Supabaseの接続URLを環境変数から取得し、コネクションプールを設定する。
// コネクションの最大数やアイドルタイム、シンプルプロトコルの使用を設定する。
// 成功時にはnilを返し、接続に失敗した場合はエラーメッセージを返す。
func SetUpSupabase() error {
	logger.InfoLog.Println("Initializing Supabase client...")

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		logger.ErrorLog.Println("DATABASE_URL is not set in environment variables")
		return fmt.Errorf("DATABASE_URL is missing")
	}
	databaseURL += "?sslmode=require"

	// コネクションプールの設定
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		logger.ErrorLog.Printf("Unable to parse database URL: %v", err)
		return fmt.Errorf("unable to parse database URL: %v", err)
	}

	// コネクションプールの設定（パフォーマンス向上のため）
	config.MaxConns = 10                                                     // 最大接続数
	config.MaxConnIdleTime = 30 * time.Second                                // アイドル接続時間
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol // Simple Protocol の使用

	// Supabaseに接続
	logger.InfoLog.Println("Connecting to Supabase database...")
	Pool, err = pgxpool.NewWithConfig(Ctx, config)
	if err != nil {
		logger.ErrorLog.Printf("Unable to connect to Supabase: %v", err)
		return fmt.Errorf("unable to connect to Supabase: %v", err)
	}

	// 接続確認
	err = Pool.Ping(Ctx)
	if err != nil {
		logger.ErrorLog.Printf("Database ping failed: %v", err)
		return fmt.Errorf("database ping failed: %v", err)
	}

	logger.InfoLog.Println("Connected to Supabase successfully")
	return nil
}

// Supabaseのコネクションプールをクローズ。
// この関数はアプリケーションのシャットダウン時に呼び出されることを想定する。
func ClosePool() {
	closeOnce.Do(func() {
		if Pool != nil {
			Pool.Close()
			logger.InfoLog.Println("Supabase connection pool closed")
		}
	})
}

// Supabaseに対してシンプルなクエリを実行し、接続が正しく動作しているかを確認する。
// クエリ結果として "1" を取得し、それをログに出力する。
// クエリに失敗した場合、エラーを返する。
func TestQuery() error {
	logger.InfoLog.Println("Testing query...")

	query := `SELECT 1`
	rows, err := Pool.Query(Ctx, query)
	if err != nil {
		logger.ErrorLog.Printf("Failed to test query: %v", err)
		return fmt.Errorf("failed to test query: %v", err)
	}
	logger.InfoLog.Println("Test query successful")
	defer rows.Close()

	for rows.Next() {
		var num int
		err := rows.Scan(&num)
		if err != nil {
			logger.ErrorLog.Printf("Failed to scan test query result: %v", err)
			return fmt.Errorf("failed to scan test query result: %v", err)
		}
		logger.InfoLog.Println("Test Query Result:", num)
	}

	logger.InfoLog.Println("Test query completed")
	return rows.Err()
}
