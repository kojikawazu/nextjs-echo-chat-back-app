package handlers_chat_likes

import (
	services_chat_likes "nextjs-echo-chat-back-app/services/chat_likes"
	services_clerk_jwt "nextjs-echo-chat-back-app/services/clerk_jwt"
)

// ChatLikesHandler は ChatLikesService のハンドラ
type ChatLikesHandler struct {
	ChatLikesService services_chat_likes.ChatLikesService
	ClerkJwtService  services_clerk_jwt.ClerkJwtService
}

// NewChatLikesHandler は ChatLikesHandler の新しいインスタンスを作成する。
func NewChatLikesHandler(chatLikesService services_chat_likes.ChatLikesService, clerkJwtService services_clerk_jwt.ClerkJwtService) *ChatLikesHandler {
	return &ChatLikesHandler{
		ChatLikesService: chatLikesService,
		ClerkJwtService:  clerkJwtService,
	}
}
