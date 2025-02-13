package handlers_chat_messages

import (
	services_chat_messages "nextjs-echo-chat-back-app/services/chat_messages"
	services_clerk_jwt "nextjs-echo-chat-back-app/services/clerk_jwt"
)

// ChatMessagesHandler は ChatMessagesService のハンドラ
type ChatMessagesHandler struct {
	ChatMessagesService services_chat_messages.ChatMessagesService
	ClerkJwtService     services_clerk_jwt.ClerkJwtService
}

// NewChatMessagesHandler は ChatMessagesHandler の新しいインスタンスを作成する。
func NewChatMessagesHandler(chatMessagesService services_chat_messages.ChatMessagesService, clerkJwtService services_clerk_jwt.ClerkJwtService) *ChatMessagesHandler {
	return &ChatMessagesHandler{
		ChatMessagesService: chatMessagesService,
		ClerkJwtService:     clerkJwtService,
	}
}
