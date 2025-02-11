package services_chat_likes

import (
	"errors"
	"nextjs-echo-chat-back-app/utils/logger"
)

// FetchChatLikesInUsers は `chat_likes` テーブルと `users` テーブルを結合して、チャットいいね情報とユーザー情報を取得する。
func (s *ChatLikesServiceImpl) FetchChatLikesInUsers(id string) ([]map[string]string, error) {
	// idが空の場合はエラー
	if id == "" {
		logger.ErrorLog.Printf("id is required")
		return nil, errors.New("id is required")
	}

	chatLikes, err := s.ChatLikesRepository.FetchChatLikesInUsers(id)
	if err != nil {
		return nil, err
	}

	return chatLikes, nil
}
