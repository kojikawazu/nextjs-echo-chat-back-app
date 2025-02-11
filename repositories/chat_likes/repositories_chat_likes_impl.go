package repositories_chat_likes

// ChatLikesRepositoryインターフェース
type ChatLikesRepository interface {
	FetchChatLikesInUsers(id string) ([]map[string]string, error)
}

// ChatLikesRepositoryImpl は ChatLikesRepository の実装
type ChatLikesRepositoryImpl struct{}

// ChatLikesRepositoryインターフェースを実装したChatLikesRepositoryImplのポインタを返す
func NewChatLikesRepository() ChatLikesRepository {
	return &ChatLikesRepositoryImpl{}
}
