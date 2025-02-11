package services_chat_messages

import (
	"errors"
	"nextjs-echo-chat-back-app/utils/logger"
	utils_uuid "nextjs-echo-chat-back-app/utils/uuid"
)

// FetchChatMessagesInRoom は `chat_messages` テーブルと `users` テーブルを結合して、チャットメッセージ情報とユーザー情報を取得する。
func (s *ChatMessagesServiceImpl) FetchChatMessagesInRoom(id string) ([]map[string]string, error) {
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

	chatMessages, err := s.ChatMessagesRepository.FetchChatMessagesInRoom(id)
	if err != nil {
		return nil, err
	}

	return chatMessages, nil
}
