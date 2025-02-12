package repositories_chat_likes

import (
	"errors"
	"nextjs-echo-chat-back-app/middlewares"
	"nextjs-echo-chat-back-app/utils/logger"
)

// FetchChatLikesInUsers は `chat_likes` テーブルと `users` テーブルを結合して、チャットいいね情報とユーザー情報を取得する。
func (r *ChatLikesRepositoryImpl) FetchChatLikesInUsers(messageId string) ([]map[string]string, error) {
	if messageId == "" {
		logger.ErrorLog.Printf("messageId is required")
		return nil, errors.New("messageId is required")
	}

	query := `
		SELECT cl.user_id, u.name, cl.created_at
		FROM chat_likes cl
		JOIN auth_users u ON cl.user_id = u.id
		WHERE cl.message_id = $1;
	`

	rows, err := middlewares.Pool.Query(middlewares.Ctx, query, messageId)
	if err != nil {
		logger.ErrorLog.Printf("Failed to fetch chat_likes: %v", err)
		return nil, err
	}

	var likes []map[string]string

	for rows.Next() {
		var userID, name, createdAt string
		err := rows.Scan(
			&userID,
			&name,
			&createdAt,
		)
		if err != nil {
			logger.ErrorLog.Printf("Failed to scan chat_likes: %v", err)
			return nil, err
		}

		likes = append(likes, map[string]string{
			"user_id":    userID,
			"name":       name,
			"created_at": createdAt,
		})
	}

	if rows.Err() != nil {
		logger.ErrorLog.Printf("Failed to fetch chat_likes: %v", rows.Err())
		return nil, rows.Err()
	}

	logger.InfoLog.Printf("Fetched %d chat_likes", len(likes))
	logger.InfoLog.Println("Fetched chat_likes successfully")
	return likes, nil
}

// CreateChatLike は `chat_likes` テーブルに新しいいいねを作成する。
func (r *ChatLikesRepositoryImpl) CreateChatLike(messageId string, userId string) (string, error) {
	if messageId == "" || userId == "" {
		logger.ErrorLog.Printf("messageId and userId are required")
		return "", errors.New("messageId and userId are required")
	}

	query := `
		INSERT INTO chat_likes (message_id, user_id)
		VALUES ($1, $2)
		RETURNING id
	`

	// トランザクションを開始
	tx, err := middlewares.Pool.Begin(middlewares.Ctx)
	if err != nil {
		logger.ErrorLog.Printf("Failed to begin transaction: %v", err)
		return "", err
	}
	// ロールバック
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback(middlewares.Ctx)
		}
	}()
	// トランザクションのロールバックを `defer` で設定（Commit された場合は無視される）
	defer tx.Rollback(middlewares.Ctx)

	// トランザクションを実行
	rows, err := tx.Query(middlewares.Ctx, query, messageId, userId)
	if err != nil {
		logger.ErrorLog.Printf("Failed to create chat_likes: %v", err)
		return "", err
	}

	// コミット
	err = tx.Commit(middlewares.Ctx)
	if err != nil {
		logger.ErrorLog.Printf("Failed to commit transaction: %v", err)
		return "", err
	}

	var likeID string
	err = rows.Scan(&likeID)
	if err != nil {
		logger.ErrorLog.Printf("Failed to scan chat_likes: %v", err)
		return "", err
	}

	logger.InfoLog.Printf("Created chat_likes successfully")
	return likeID, nil
}

// DeleteChatLike は `chat_likes` テーブルからいいねを削除する。
func (r *ChatLikesRepositoryImpl) DeleteChatLike(messageId string, userId string) (string, error) {
	if messageId == "" || userId == "" {
		logger.ErrorLog.Printf("messageId and userId are required")
		return "", errors.New("messageId and userId are required")
	}

	query := `
		DELETE FROM chat_likes
		WHERE message_id = $1 AND user_id = $2;
	`

	// トランザクションを開始
	tx, err := middlewares.Pool.Begin(middlewares.Ctx)
	if err != nil {
		logger.ErrorLog.Printf("Failed to begin transaction: %v", err)
		return "", err
	}
	// ロールバック
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback(middlewares.Ctx)
		}
	}()
	// トランザクションのロールバックを `defer` で設定（Commit された場合は無視される）
	defer tx.Rollback(middlewares.Ctx)

	// トランザクションを実行
	_, err = tx.Query(middlewares.Ctx, query, messageId, userId)
	if err != nil {
		logger.ErrorLog.Printf("Failed to delete chat_likes: %v", err)
		return "", err
	}

	// コミット
	err = tx.Commit(middlewares.Ctx)
	if err != nil {
		logger.ErrorLog.Printf("Failed to commit transaction: %v", err)
		return "", err
	}

	logger.InfoLog.Printf("Deleted chat_likes successfully")
	return "Deleted chat_likes successfully", nil
}
