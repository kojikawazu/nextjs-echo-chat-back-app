package services_chat_messages

import (
	"errors"
	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"
	utils_uuid "nextjs-echo-chat-back-app/utils/uuid"
)

// FetchChatMessagesInRoom は `chat_messages` テーブルと `users` テーブルを結合して、チャットメッセージ情報とユーザー情報を取得する。
func (s *ChatMessagesServiceImpl) FetchChatMessagesInRoom(roomId string) ([]map[string]interface{}, error) {
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

	chatMessages, err := s.ChatMessagesRepository.FetchChatMessagesInRoom(roomId)
	if err != nil {
		return nil, err
	}

	return chatMessages, nil
}

// CreateChatMessage は `chat_messages` テーブルにメッセージを作成する。
func (s *ChatMessagesServiceImpl) CreateChatMessage(message string, roomId string, userId string) (models.ChatMessages, error) {
	// idが空の場合はエラー
	if message == "" || roomId == "" || userId == "" {
		logger.ErrorLog.Printf("message, roomId, userId is required")
		return models.ChatMessages{}, errors.New("message, roomId, userId is required")
	}
	// UUIDかどうかを確認
	if !utils_uuid.IsUUID(roomId) {
		logger.ErrorLog.Printf("invalid id")
		return models.ChatMessages{}, errors.New("invalid id")
	}

	// メッセージを作成
	msg, err := s.ChatMessagesRepository.CreateChatMessage(message, roomId, userId)
	if err != nil {
		return models.ChatMessages{}, err
	}

	logger.InfoLog.Printf("Created chat_message successfully")
	return msg, nil
}
