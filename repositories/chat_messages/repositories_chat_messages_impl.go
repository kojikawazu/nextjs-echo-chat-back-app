package repositories_chat_messages

// ChatMessagesRepositoryインターフェース
type ChatMessagesRepository interface {
	FetchChatMessagesInRoom(id string) ([]map[string]string, error)
}

// ChatMessagesRepositoryImpl は ChatMessagesRepository の実装
type ChatMessagesRepositoryImpl struct{}

// ChatMessagesRepositoryインターフェースを実装したChatMessagesRepositoryImplのポインタを返す
func NewChatMessagesRepository() ChatMessagesRepository {
	return &ChatMessagesRepositoryImpl{}
}
