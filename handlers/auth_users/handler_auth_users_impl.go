package handlers_auth_users

import (
	services_auth_users "nextjs-echo-chat-back-app/services/auth_users"
	services_clerk_jwt "nextjs-echo-chat-back-app/services/clerk_jwt"
)

// AuthUsersHandler は AuthUsersService のハンドラ
type AuthUsersHandler struct {
	AuthUsersService services_auth_users.AuthUsersService
	ClerkJwtService  services_clerk_jwt.ClerkJwtService
}

// NewAuthUsersHandler は AuthUsersHandler の新しいインスタンスを作成する。
func NewAuthUsersHandler(authUsersService services_auth_users.AuthUsersService, clerkJwtService services_clerk_jwt.ClerkJwtService) *AuthUsersHandler {
	return &AuthUsersHandler{
		AuthUsersService: authUsersService,
		ClerkJwtService:  clerkJwtService,
	}
}
