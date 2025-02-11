package repositories_chat_messages

import (
	"encoding/json"
	"errors"
	"nextjs-echo-chat-back-app/middlewares"
	"nextjs-echo-chat-back-app/utils/logger"
)

// FetchChatMessagesInRoom は `chat_messages` テーブルと `users` テーブルを結合して、チャットメッセージ情報とユーザー情報を取得する。
func (r *ChatMessagesRepositoryImpl) FetchChatMessagesInRoom(roomId string) ([]map[string]interface{}, error) {
	if roomId == "" {
		logger.ErrorLog.Printf("id is required")
		return nil, errors.New("id is required")
	}

	query := `
		SELECT 
			cm.id AS message_id, 
			cm.user_id, 
			u.name, 
			cm.content, 
			cm.created_at,
			COUNT(cl.id) AS like_count,
			COALESCE(
				json_agg(DISTINCT jsonb_build_object('user_id', cl.user_id, 'name', lu.name))
				FILTER (WHERE cl.user_id IS NOT NULL), '[]'
			) AS liked_users
		FROM chat_messages cm
		JOIN auth_users u ON cm.user_id = u.id
		LEFT JOIN chat_likes cl ON cm.id = cl.message_id
		LEFT JOIN auth_users lu ON cl.user_id = lu.id
		WHERE cm.room_id = $1
		GROUP BY cm.id, cm.user_id, u.name, cm.content, cm.created_at
		ORDER BY cm.created_at ASC;
	`
	rows, err := middlewares.Pool.Query(middlewares.Ctx, query, roomId)
	if err != nil {
		logger.ErrorLog.Printf("Failed to fetch chat_messages: %v", err)
		return nil, err
	}

	var messages []map[string]interface{}

	for rows.Next() {
		var messageID, userID, name, content, createdAt string
		var likeCount int
		var likedUsersJSON string
		err := rows.Scan(
			&messageID,
			&userID,
			&name,
			&content,
			&createdAt,
			&likeCount,
			&likedUsersJSON,
		)

		if err != nil {
			logger.ErrorLog.Printf("Failed to scan chat_messages: %v", err)
			return nil, err
		}

		// `likedUsersJSON` をパース
		var likedUsers []map[string]string
		err = json.Unmarshal([]byte(likedUsersJSON), &likedUsers)
		if err != nil {
			logger.ErrorLog.Printf("Failed to parse liked_users JSON: %v", err)
			return nil, err
		}

		messages = append(messages, map[string]interface{}{
			"message_id":  messageID,
			"user_id":     userID,
			"user_name":   name,
			"content":     content,
			"created_at":  createdAt,
			"like_count":  likeCount,
			"liked_users": likedUsers,
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
