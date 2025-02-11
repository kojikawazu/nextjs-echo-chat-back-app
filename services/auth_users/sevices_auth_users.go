package services_auth_users

import (
	"nextjs-echo-chat-back-app/models"
)

// AuthUsersService は `auth_users` テーブルからすべてのユーザー情報を取得する。
func (r *AuthUsersServiceImpl) FetchAuthUsers() ([]models.AuthUsers, error) {
	authUsers, err := r.AuthUsersRepository.FetchAuthUsers()
	if err != nil {
		return nil, err
	}
	return authUsers, nil
}
