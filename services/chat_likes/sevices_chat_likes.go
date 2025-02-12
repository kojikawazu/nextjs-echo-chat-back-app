package services_chat_likes

import (
	"errors"
	"nextjs-echo-chat-back-app/utils/logger"
	utils_uuid "nextjs-echo-chat-back-app/utils/uuid"
)

// FetchChatLikesInUsers は `chat_likes` テーブルと `users` テーブルを結合して、チャットいいね情報とユーザー情報を取得する。
func (s *ChatLikesServiceImpl) FetchChatLikesInUsers(messageId string) ([]map[string]string, error) {
	// messageIdが空の場合はエラー
	if messageId == "" {
		logger.ErrorLog.Printf("messageId is required")
		return nil, errors.New("messageId is required")
	}

	chatLikes, err := s.ChatLikesRepository.FetchChatLikesInUsers(messageId)
	if err != nil {
		return nil, err
	}

	return chatLikes, nil
}

// CreateChatLike は `chat_likes` テーブルに新しいいいねを作成する。
func (s *ChatLikesServiceImpl) CreateChatLike(messageId string, userId string) (string, error) {
	// messageIdが空の場合はエラー
	if messageId == "" {
		logger.ErrorLog.Printf("messageId is required")
		return "", errors.New("messageId is required")
	}
	// userIdが空の場合はエラー
	if userId == "" {
		logger.ErrorLog.Printf("userId is required")
		return "", errors.New("userId is required")
	}
	// UUIDかどうかを確認
	if !utils_uuid.IsUUID(messageId) {
		logger.ErrorLog.Printf("invalid messageId")
		return "", errors.New("invalid messageId")
	}

	// いいねを作成
	likeId, err := s.ChatLikesRepository.CreateChatLike(messageId, userId)
	if err != nil {
		return "", err
	}

	return likeId, nil
}

// DeleteChatLike は `chat_likes` テーブルからいいねを削除する。
func (s *ChatLikesServiceImpl) DeleteChatLike(messageId string, userId string) (string, error) {
	// messageIdが空の場合はエラー
	if messageId == "" {
		logger.ErrorLog.Printf("messageId is required")
		return "", errors.New("messageId is required")
	}
	// userIdが空の場合はエラー
	if userId == "" {
		logger.ErrorLog.Printf("userId is required")
		return "", errors.New("userId is required")
	}
	// UUIDかどうかを確認
	if !utils_uuid.IsUUID(messageId) {
		logger.ErrorLog.Printf("invalid messageId")
		return "", errors.New("invalid messageId")
	}

	// いいねを削除
	likeId, err := s.ChatLikesRepository.DeleteChatLike(messageId, userId)
	if err != nil {
		return "", err
	}

	return likeId, nil
}
