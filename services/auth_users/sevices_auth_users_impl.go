package services_auth_users

import (
	"nextjs-echo-chat-back-app/models"
	repositories_auth_users "nextjs-echo-chat-back-app/repositories/auth_users"
)

// AuthUsersServiceインターフェース
type AuthUsersService interface {
	FetchAuthUsers() ([]models.AuthUsers, error)
}

// AuthUsersServiceImpl は AuthUsersService の実装
type AuthUsersServiceImpl struct {
	AuthUsersRepository repositories_auth_users.AuthUsersRepository
}

// AuthUsersServiceインターフェースを実装したAuthUsersServiceImplのポインタを返す
func NewAuthUsersService(
	authUsersRepository repositories_auth_users.AuthUsersRepository,
) AuthUsersService {
	return &AuthUsersServiceImpl{
		AuthUsersRepository: authUsersRepository,
	}
}
