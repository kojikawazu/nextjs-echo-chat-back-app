package repositories_chat_messages

// ChatMessagesRepositoryインターフェース
type ChatMessagesRepository interface {
	FetchChatMessagesInRoom(roomId string) ([]map[string]interface{}, error)
	CreateChatMessage(message string, roomId string, userId string) (string, error)
}

// ChatMessagesRepositoryImpl は ChatMessagesRepository の実装
type ChatMessagesRepositoryImpl struct{}

// ChatMessagesRepositoryインターフェースを実装したChatMessagesRepositoryImplのポインタを返す
func NewChatMessagesRepository() ChatMessagesRepository {
	return &ChatMessagesRepositoryImpl{}
}
