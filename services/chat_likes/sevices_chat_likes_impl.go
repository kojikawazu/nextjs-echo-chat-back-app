package services_chat_likes

import (
	repositories_chat_likes "nextjs-echo-chat-back-app/repositories/chat_likes"
)

// ChatLikesServiceインターフェース
type ChatLikesService interface {
	FetchChatLikesInUsers(encryptedMessageId string) ([]map[string]string, error)
	CreateChatLike(encryptedMessageId string, userId string) (string, error)
	DeleteChatLike(encryptedMessageId string, userId string) (string, error)
}

// ChatMessagesServiceImpl は ChatMessagesService の実装
type ChatLikesServiceImpl struct {
	ChatLikesRepository repositories_chat_likes.ChatLikesRepository
}

// NewChatMessagesService は ChatMessagesService の新しいインスタンスを作成する。
func NewChatLikesService(chatLikesRepository repositories_chat_likes.ChatLikesRepository) ChatLikesService {
	return &ChatLikesServiceImpl{ChatLikesRepository: chatLikesRepository}
}
