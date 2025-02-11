package handlers_chat_messages

import services_chat_messages "nextjs-echo-chat-back-app/services/chat_messages"

// ChatMessagesHandler は ChatMessagesService のハンドラ
type ChatMessagesHandler struct {
	ChatMessagesService services_chat_messages.ChatMessagesService
}

// NewChatMessagesHandler は ChatMessagesHandler の新しいインスタンスを作成する。
func NewChatMessagesHandler(chatMessagesService services_chat_messages.ChatMessagesService) *ChatMessagesHandler {
	return &ChatMessagesHandler{
		ChatMessagesService: chatMessagesService,
	}
}
