package services_chat_messages

import (
	repositories_chat_messages "nextjs-echo-chat-back-app/repositories/chat_messages"
)

// ChatMessagesServiceインターフェース
type ChatMessagesService interface {
	FetchChatMessagesInRoom(id string) ([]map[string]string, error)
}

// ChatMessagesServiceImpl は ChatMessagesService の実装
type ChatMessagesServiceImpl struct {
	ChatMessagesRepository repositories_chat_messages.ChatMessagesRepository
}

// NewChatMessagesService は ChatMessagesService の新しいインスタンスを作成する。
func NewChatMessagesService(chatMessagesRepository repositories_chat_messages.ChatMessagesRepository) ChatMessagesService {
	return &ChatMessagesServiceImpl{ChatMessagesRepository: chatMessagesRepository}
}
