package repositories_auth_users

import "nextjs-echo-chat-back-app/models"

// AuthUsersRepositoryインターフェース
type AuthUsersRepository interface {
	FetchAuthUsers() ([]models.AuthUsers, error)
}

// AuthUsersRepositoryImpl は AuthUsersRepository の実装
type AuthUsersRepositoryImpl struct{}

// AuthUsersRepositoryインターフェースを実装したAuthUsersRepositoryImplのポインタを返す
func NewAuthUsersRepository() AuthUsersRepository {
	return &AuthUsersRepositoryImpl{}
}
