package repositories_chat_rooms

import (
	"encoding/json"
	"errors"
	"nextjs-echo-chat-back-app/middlewares"
	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"
)

// FetchChatRooms は `chat_rooms` テーブルからすべてのチャットルーム情報を取得する。
func (r *ChatRoomsRepositoryImpl) FetchChatRooms() ([]models.ChatRooms, error) {
	query := `
		SELECT r.id, r.name, r.created_at, r.updated_at,
		       COALESCE(json_agg(json_build_object('id', u.id, 'name', u.name)) 
		                FILTER (WHERE u.id IS NOT NULL), '[]') AS users
		FROM chat_rooms r
		LEFT JOIN rooms_users ru ON r.id = ru.room_id
		LEFT JOIN auth_users u ON ru.user_id = u.id
		GROUP BY r.id;
	`

	rows, err := middlewares.Pool.Query(middlewares.Ctx, query)
	if err != nil {
		logger.ErrorLog.Printf("Failed to fetch chat_rooms: %v", err)
		return nil, err
	}
	defer rows.Close()

	var chatRooms []models.ChatRooms
	for rows.Next() {
		var chatRoom models.ChatRooms
		var usersJSON string

		err := rows.Scan(
			&chatRoom.ID,
			&chatRoom.Name,
			&chatRoom.CreatedAt,
			&chatRoom.UpdatedAt,
			&usersJSON,
		)
		if err != nil {
			logger.ErrorLog.Printf("Failed to scan chat_rooms: %v", err)
			return nil, err
		}

		// usersJSON をパースして []map[string]string に変換
		var users []models.MiniAuthUsers
		err = json.Unmarshal([]byte(usersJSON), &users)
		if err != nil {
			logger.ErrorLog.Printf("Failed to parse users JSON: %v", err)
			return nil, err
		}
		chatRoom.Users = users

		chatRooms = append(chatRooms, chatRoom)
	}

	if rows.Err() != nil {
		logger.ErrorLog.Printf("Failed to fetch chat_rooms: %v", rows.Err())
		return nil, rows.Err()
	}

	logger.InfoLog.Printf("Fetched %d chat_rooms", len(chatRooms))
	logger.InfoLog.Println("Fetched chat_rooms successfully")
	return chatRooms, nil
}

// FetchUsersInRoom は `chat_rooms` テーブルと `users` テーブルを結合して、チャットルーム情報とユーザー情報を取得する。
func (r *ChatRoomsRepositoryImpl) FetchUsersInRoom(roomId string) ([]map[string]string, error) {
	if roomId == "" {
		return nil, errors.New("id is required")
	}

	query := `
		SELECT ru.user_id, u.name
		FROM rooms_users ru
		JOIN auth_users u ON ru.user_id = u.id
		WHERE ru.room_id = $1
	`
	rows, err := middlewares.Pool.Query(middlewares.Ctx, query, roomId)
	if err != nil {
		logger.ErrorLog.Printf("Failed to fetch chat_rooms: %v", err)
		return nil, err
	}
	defer rows.Close()
	var users []map[string]string

	for rows.Next() {
		var userID, name string
		err := rows.Scan(
			&userID,
			&name,
		)

		if err != nil {
			logger.ErrorLog.Printf("Failed to scan room_users: %v", err)
			return nil, err
		}

		users = append(users, map[string]string{
			"user_id": userID,
			"name":    name,
		})
	}

	if rows.Err() != nil {
		logger.ErrorLog.Printf("Failed to fetch chat_rooms: %v", rows.Err())
		return nil, rows.Err()
	}

	logger.InfoLog.Printf("Fetched %d users", len(users))
	logger.InfoLog.Println("Fetched users successfully")
	return users, nil
}

// CreateRoom は `chat_rooms` テーブルに新しいチャットルームを作成する。
func (r *ChatRoomsRepositoryImpl) CreateRoom(roomName string) (string, error) {
	if roomName == "" {
		return "", errors.New("room name is required")
	}

	query := `
		INSERT INTO chat_rooms (name) 
		VALUES ($1) 
		RETURNING id
	`

	// トランザクション開始
	tx, err := middlewares.Pool.Begin(middlewares.Ctx)
	if err != nil {
		logger.ErrorLog.Printf("Failed to begin transaction: %v", err)
		return "", err
	}
	defer func() {
		// ロールバック
		if r := recover(); r != nil {
			tx.Rollback(middlewares.Ctx)
		}
	}()
	// トランザクションのロールバックを `defer` で設定（Commit された場合は無視される）
	defer tx.Rollback(middlewares.Ctx)

	// `QueryRow` を使用して 1 つの値を取得
	var roomID string
	err = tx.QueryRow(middlewares.Ctx, query, roomName).Scan(&roomID)
	if err != nil {
		logger.ErrorLog.Printf("Failed to create room: %v", err)
		return "", err
	}

	// コミット
	if err := tx.Commit(middlewares.Ctx); err != nil {
		logger.ErrorLog.Printf("Failed to commit transaction: %v", err)
		return "", err
	}

	logger.InfoLog.Printf("Created room with ID: %s", roomID)
	return roomID, nil
}
