package repositories_chat_rooms

import "nextjs-echo-chat-back-app/models"

// ChatRoomsRepositoryインターフェース
type ChatRoomsRepository interface {
	FetchChatRooms() ([]models.ChatRooms, error)
	FetchUsersInRoom(id string) ([]map[string]string, error)
}

// ChatRoomsRepositoryImpl は ChatRoomsRepository の実装
type ChatRoomsRepositoryImpl struct{}

// ChatRoomsRepositoryインターフェースを実装したChatRoomsRepositoryImplのポインタを返す
func NewChatRoomsRepository() ChatRoomsRepository {
	return &ChatRoomsRepositoryImpl{}
}
