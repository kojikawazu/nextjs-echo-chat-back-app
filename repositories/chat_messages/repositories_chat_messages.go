package repositories_chat_messages

import (
	"errors"
	"nextjs-echo-chat-back-app/middlewares"
	"nextjs-echo-chat-back-app/utils/logger"
)

// FetchChatMessagesInRoom は `chat_messages` テーブルと `users` テーブルを結合して、チャットメッセージ情報とユーザー情報を取得する。
func (r *ChatMessagesRepositoryImpl) FetchChatMessagesInRoom(id string) ([]map[string]string, error) {
	if id == "" {
		logger.ErrorLog.Printf("id is required")
		return nil, errors.New("id is required")
	}

	query := `
		SELECT cm.id, cm.user_id, u.name, cm.content, cm.created_at
		FROM chat_messages cm
		JOIN auth_users u ON cm.user_id = u.id
		WHERE cm.room_id = $1
		ORDER BY cm.created_at ASC;
	`
	rows, err := middlewares.Pool.Query(middlewares.Ctx, query, id)
	if err != nil {
		logger.ErrorLog.Printf("Failed to fetch chat_messages: %v", err)
		return nil, err
	}

	var messages []map[string]string

	for rows.Next() {
		var messageID, userID, name, content, createdAt string
		err := rows.Scan(
			&messageID,
			&userID,
			&name,
			&content,
			&createdAt,
		)

		if err != nil {
			logger.ErrorLog.Printf("Failed to scan chat_messages: %v", err)
			return nil, err
		}

		messages = append(messages, map[string]string{
			"message_id": messageID,
			"user_id":    userID,
			"name":       name,
			"content":    content,
			"created_at": createdAt,
		})
	}

	if rows.Err() != nil {
		logger.ErrorLog.Printf("Failed to fetch chat_messages: %v", rows.Err())
		return nil, rows.Err()
	}

	logger.InfoLog.Printf("Fetched %d messages", len(messages))
	logger.InfoLog.Println("Fetched messages successfully")
	return messages, nil
}
