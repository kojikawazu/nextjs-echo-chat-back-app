package handlers_chat_likes

import services_chat_likes "nextjs-echo-chat-back-app/services/chat_likes"

// ChatLikesHandler は ChatLikesService のハンドラ
type ChatLikesHandler struct {
	ChatLikesService services_chat_likes.ChatLikesService
}

// NewChatLikesHandler は ChatLikesHandler の新しいインスタンスを作成する。
func NewChatLikesHandler(chatLikesService services_chat_likes.ChatLikesService) *ChatLikesHandler {
	return &ChatLikesHandler{
		ChatLikesService: chatLikesService,
	}
}
