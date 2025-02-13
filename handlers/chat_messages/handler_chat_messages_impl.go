package handlers_chat_messages

import (
	handlers_websocket_messages "nextjs-echo-chat-back-app/handlers/websocket_messages"
	services_chat_messages "nextjs-echo-chat-back-app/services/chat_messages"
	services_clerk_jwt "nextjs-echo-chat-back-app/services/clerk_jwt"
)

// ChatMessagesHandler は ChatMessagesService のハンドラ
type ChatMessagesHandler struct {
	ChatMessagesService services_chat_messages.ChatMessagesService
	ClerkJwtService     services_clerk_jwt.ClerkJwtService
	WebSocketHandler    *handlers_websocket_messages.WebSocketHandler
}

// NewChatMessagesHandler は ChatMessagesHandler の新しいインスタンスを作成する。
func NewChatMessagesHandler(
	chatMessagesService services_chat_messages.ChatMessagesService,
	clerkJwtService services_clerk_jwt.ClerkJwtService,
	webSocketHandler *handlers_websocket_messages.WebSocketHandler,
) *ChatMessagesHandler {
	return &ChatMessagesHandler{
		ChatMessagesService: chatMessagesService,
		ClerkJwtService:     clerkJwtService,
		WebSocketHandler:    webSocketHandler,
	}
}
