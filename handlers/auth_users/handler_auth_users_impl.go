package handlers_auth_users

import services_auth_users "nextjs-echo-chat-back-app/services/auth_users"

// AuthUsersHandler は AuthUsersService のハンドラ
type AuthUsersHandler struct {
	AuthUsersService services_auth_users.AuthUsersService
}

// NewAuthUsersHandler は AuthUsersHandler の新しいインスタンスを作成する。
func NewAuthUsersHandler(authUsersService services_auth_users.AuthUsersService) *AuthUsersHandler {
	return &AuthUsersHandler{
		AuthUsersService: authUsersService,
	}
}
