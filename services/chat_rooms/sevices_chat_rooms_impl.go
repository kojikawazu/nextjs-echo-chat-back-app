package services_chat_rooms

import (
	"nextjs-echo-chat-back-app/models"
	repositories_chat_rooms "nextjs-echo-chat-back-app/repositories/chat_rooms"
)

// ChatRoomsServiceインターフェース
type ChatRoomsService interface {
	FetchChatRooms() ([]models.ChatRooms, error)
	FetchUsersInRoom(id string) ([]map[string]string, error)
}

// ChatRoomsServiceImpl は ChatRoomsService の実装
type ChatRoomsServiceImpl struct {
	ChatRoomsRepository repositories_chat_rooms.ChatRoomsRepository
}

// NewChatRoomsService は ChatRoomsService の新しいインスタンスを作成する。
func NewChatRoomsService(chatRoomsRepository repositories_chat_rooms.ChatRoomsRepository) ChatRoomsService {
	return &ChatRoomsServiceImpl{ChatRoomsRepository: chatRoomsRepository}
}
