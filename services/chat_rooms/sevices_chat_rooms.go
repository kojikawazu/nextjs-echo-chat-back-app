package services_chat_rooms

import (
	"errors"
	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"
	utils_uuid "nextjs-echo-chat-back-app/utils/uuid"
)

// FetchChatRooms は `chat_rooms` テーブルからすべてのチャットルーム情報を取得する。
func (s *ChatRoomsServiceImpl) FetchChatRooms() ([]models.ChatRooms, error) {
	chatRooms, err := s.ChatRoomsRepository.FetchChatRooms()
	if err != nil {
		return nil, err
	}
	return chatRooms, nil
}

// FetchUsersInRoom は `chat_rooms` テーブルと `users` テーブルを結合して、チャットルーム情報とユーザー情報を取得する。
func (s *ChatRoomsServiceImpl) FetchUsersInRoom(id string) ([]map[string]string, error) {
	// idが空の場合はエラー
	if id == "" {
		logger.ErrorLog.Printf("id is required")
		return nil, errors.New("id is required")
	}
	// UUIDかどうかを確認
	if !utils_uuid.IsUUID(id) {
		logger.ErrorLog.Printf("invalid id")
		return nil, errors.New("invalid id")
	}

	users, err := s.ChatRoomsRepository.FetchUsersInRoom(id)
	if err != nil {
		return nil, err
	}

	return users, nil
}
