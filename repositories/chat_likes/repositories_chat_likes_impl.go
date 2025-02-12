package repositories_chat_likes

// ChatLikesRepositoryインターフェース
type ChatLikesRepository interface {
	FetchChatLikesInUsers(messageId string) ([]map[string]string, error)
	CreateChatLike(messageId string, userId string) (string, error)
	DeleteChatLike(messageId string, userId string) (string, error)
}

// ChatLikesRepositoryImpl は ChatLikesRepository の実装
type ChatLikesRepositoryImpl struct{}

// ChatLikesRepositoryインターフェースを実装したChatLikesRepositoryImplのポインタを返す
func NewChatLikesRepository() ChatLikesRepository {
	return &ChatLikesRepositoryImpl{}
}
