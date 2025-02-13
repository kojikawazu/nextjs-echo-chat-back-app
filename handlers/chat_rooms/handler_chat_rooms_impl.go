package handlers_chat_rooms

import (
	services_chat_rooms "nextjs-echo-chat-back-app/services/chat_rooms"
	services_clerk_jwt "nextjs-echo-chat-back-app/services/clerk_jwt"
)

// ChatRoomsHandler は ChatRoomsService のハンドラ
type ChatRoomsHandler struct {
	ChatRoomsService services_chat_rooms.ChatRoomsService
	ClerkJwtService  services_clerk_jwt.ClerkJwtService
}

// NewChatRoomsHandler は ChatRoomsHandler の新しいインスタンスを作成する。
func NewChatRoomsHandler(chatRoomsService services_chat_rooms.ChatRoomsService, clerkJwtService services_clerk_jwt.ClerkJwtService) *ChatRoomsHandler {
	return &ChatRoomsHandler{
		ChatRoomsService: chatRoomsService,
		ClerkJwtService:  clerkJwtService,
	}
}
