package routes

import (
	"net/http"

	handlers_auth_users "nextjs-echo-chat-back-app/handlers/auth_users"
	handlers_chat_likes "nextjs-echo-chat-back-app/handlers/chat_likes"
	handlers_chat_messages "nextjs-echo-chat-back-app/handlers/chat_messages"
	handlers_chat_rooms "nextjs-echo-chat-back-app/handlers/chat_rooms"
	repositories_auth_users "nextjs-echo-chat-back-app/repositories/auth_users"
	repositories_chat_likes "nextjs-echo-chat-back-app/repositories/chat_likes"
	repositories_chat_messages "nextjs-echo-chat-back-app/repositories/chat_messages"
	repositories_chat_rooms "nextjs-echo-chat-back-app/repositories/chat_rooms"
	services_auth_users "nextjs-echo-chat-back-app/services/auth_users"
	services_chat_likes "nextjs-echo-chat-back-app/services/chat_likes"
	services_chat_messages "nextjs-echo-chat-back-app/services/chat_messages"
	services_chat_rooms "nextjs-echo-chat-back-app/services/chat_rooms"

	"github.com/labstack/echo"
)

func SetUpRouter(e *echo.Echo) {
	// ヘルスチェックエンドポイントの追加
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service is running")
	})

	// DI注入
	// repositories
	AuthUsersRepository := repositories_auth_users.NewAuthUsersRepository()
	ChatRoomsRepository := repositories_chat_rooms.NewChatRoomsRepository()
	ChatMessagesRepository := repositories_chat_messages.NewChatMessagesRepository()
	ChatLikesRepository := repositories_chat_likes.NewChatLikesRepository()
	// services
	AuthUsersService := services_auth_users.NewAuthUsersService(AuthUsersRepository)
	ChatRoomsService := services_chat_rooms.NewChatRoomsService(ChatRoomsRepository)
	ChatMessagesService := services_chat_messages.NewChatMessagesService(ChatMessagesRepository)
	ChatLikesService := services_chat_likes.NewChatLikesService(ChatLikesRepository)
	// handlers
	AuthUsersHandler := handlers_auth_users.NewAuthUsersHandler(AuthUsersService)
	ChatRoomsHandler := handlers_chat_rooms.NewChatRoomsHandler(ChatRoomsService)
	ChatMessagesHandler := handlers_chat_messages.NewChatMessagesHandler(ChatMessagesService)
	ChatLikesHandler := handlers_chat_likes.NewChatLikesHandler(ChatLikesService)

	api := e.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("", AuthUsersHandler.FetchAuthUsers)
		}
		rooms := api.Group("/rooms")
		{
			rooms.GET("", ChatRoomsHandler.FetchChatRooms)
			rooms.GET("/:id/users", ChatRoomsHandler.FetchUsersInRoom)
			rooms.GET("/:id/messages", ChatMessagesHandler.FetchChatMessagesInRoom)
		}
		likes := api.Group("/messages")
		{
			likes.GET("/:id/likes", ChatLikesHandler.FetchChatLikesInUsers)
		}
	}
}
