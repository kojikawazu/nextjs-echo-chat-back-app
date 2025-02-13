package repositories_chat_messages

import "nextjs-echo-chat-back-app/models"

// ChatMessagesRepositoryインターフェース
type ChatMessagesRepository interface {
	GetChatMessageByID(messageId string) (models.ChatMessages, error)
	FetchChatMessagesInRoom(roomId string) ([]map[string]interface{}, error)
	CreateChatMessage(message string, roomId string, userId string) (models.ChatMessages, error)
}

// ChatMessagesRepositoryImpl は ChatMessagesRepository の実装
type ChatMessagesRepositoryImpl struct{}

// ChatMessagesRepositoryインターフェースを実装したChatMessagesRepositoryImplのポインタを返す
func NewChatMessagesRepository() ChatMessagesRepository {
	return &ChatMessagesRepositoryImpl{}
}
