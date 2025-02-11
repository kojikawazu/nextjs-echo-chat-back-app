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
func (s *ChatRoomsServiceImpl) FetchUsersInRoom(roomId string) ([]map[string]string, error) {
	// idが空の場合はエラー
	if roomId == "" {
		logger.ErrorLog.Printf("id is required")
		return nil, errors.New("id is required")
	}
	// UUIDかどうかを確認
	if !utils_uuid.IsUUID(roomId) {
		logger.ErrorLog.Printf("invalid id")
		return nil, errors.New("invalid id")
	}

	users, err := s.ChatRoomsRepository.FetchUsersInRoom(roomId)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// CreateRoom は `chat_rooms` テーブルに新しいチャットルームを作成する。
func (s *ChatRoomsServiceImpl) CreateRoom(roomName string) (string, error) {
	// ルーム名が空の場合はエラー
	if roomName == "" {
		logger.ErrorLog.Printf("room name is required")
		return "", errors.New("room name is required")
	}

	roomId, err := s.ChatRoomsRepository.CreateRoom(roomName)
	if err != nil {
		return "", err
	}

	return roomId, nil
}
