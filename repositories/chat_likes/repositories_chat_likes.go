package repositories_chat_likes

import (
	"errors"
	"nextjs-echo-chat-back-app/middlewares"
	"nextjs-echo-chat-back-app/utils/logger"
)

// FetchChatLikesInUsers は `chat_likes` テーブルと `users` テーブルを結合して、チャットいいね情報とユーザー情報を取得する。
func (r *ChatLikesRepositoryImpl) FetchChatLikesInUsers(id string) ([]map[string]string, error) {
	if id == "" {
		logger.ErrorLog.Printf("id is required")
		return nil, errors.New("id is required")
	}

	query := `
		SELECT cl.user_id, u.name, cl.created_at
		FROM chat_likes cl
		JOIN auth_users u ON cl.user_id = u.id
		WHERE cl.message_id = $1;
	`

	rows, err := middlewares.Pool.Query(middlewares.Ctx, query, id)
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
