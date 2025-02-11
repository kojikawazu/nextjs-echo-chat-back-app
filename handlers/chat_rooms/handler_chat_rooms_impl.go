package handlers_chat_rooms

import services_chat_rooms "nextjs-echo-chat-back-app/services/chat_rooms"

// ChatRoomsHandler は ChatRoomsService のハンドラ
type ChatRoomsHandler struct {
	ChatRoomsService services_chat_rooms.ChatRoomsService
}

// NewChatRoomsHandler は ChatRoomsHandler の新しいインスタンスを作成する。
func NewChatRoomsHandler(chatRoomsService services_chat_rooms.ChatRoomsService) *ChatRoomsHandler {
	return &ChatRoomsHandler{
		ChatRoomsService: chatRoomsService,
	}
}
