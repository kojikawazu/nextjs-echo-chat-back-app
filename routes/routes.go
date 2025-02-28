package routes

import (
	handlers_chat_likes "nextjs-echo-chat-back-app/handlers/chat_likes"
	handlers_chat_messages "nextjs-echo-chat-back-app/handlers/chat_messages"
	handlers_chat_rooms "nextjs-echo-chat-back-app/handlers/chat_rooms"
	handlers_websocket_messages "nextjs-echo-chat-back-app/handlers/websocket_messages"
	repositories_chat_likes "nextjs-echo-chat-back-app/repositories/chat_likes"
	repositories_chat_messages "nextjs-echo-chat-back-app/repositories/chat_messages"
	repositories_chat_rooms "nextjs-echo-chat-back-app/repositories/chat_rooms"
	services_chat_likes "nextjs-echo-chat-back-app/services/chat_likes"
	services_chat_messages "nextjs-echo-chat-back-app/services/chat_messages"
	services_chat_rooms "nextjs-echo-chat-back-app/services/chat_rooms"
	services_clerk_jwt "nextjs-echo-chat-back-app/services/clerk_jwt"

	"github.com/labstack/echo"
)

func SetUpRouter(e *echo.Echo) {
	// ヘルスチェックエンドポイントの追加
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Service is running")
	// })

	// DI注入
	// repositories
	//AuthUsersRepository := repositories_auth_users.NewAuthUsersRepository()
	ChatRoomsRepository := repositories_chat_rooms.NewChatRoomsRepository()
	ChatMessagesRepository := repositories_chat_messages.NewChatMessagesRepository()
	ChatLikesRepository := repositories_chat_likes.NewChatLikesRepository()
	// services
	//AuthUsersService := services_auth_users.NewAuthUsersService(AuthUsersRepository)
	ChatRoomsService := services_chat_rooms.NewChatRoomsService(ChatRoomsRepository)
	ChatMessagesService := services_chat_messages.NewChatMessagesService(ChatMessagesRepository)
	ChatLikesService := services_chat_likes.NewChatLikesService(ChatLikesRepository)
	ClerkJwtService := services_clerk_jwt.NewClerkJwtService()
	// handlers
	WebSocketHandler := handlers_websocket_messages.NewWebSocketHandler()
	//AuthUsersHandler := handlers_auth_users.NewAuthUsersHandler(AuthUsersService, ClerkJwtService)
	ChatRoomsHandler := handlers_chat_rooms.NewChatRoomsHandler(ChatRoomsService, ClerkJwtService)
	ChatMessagesHandler := handlers_chat_messages.NewChatMessagesHandler(ChatMessagesService, ClerkJwtService, WebSocketHandler)
	ChatLikesHandler := handlers_chat_likes.NewChatLikesHandler(ChatLikesService, ClerkJwtService)

	api := e.Group("/api")
	{
		// users := api.Group("/users")
		// {
		// 	users.GET("", AuthUsersHandler.FetchAuthUsers)
		// }
		rooms := api.Group("/rooms")
		{
			rooms.GET("", ChatRoomsHandler.FetchChatRooms)
			//rooms.GET("/:id/users", ChatRoomsHandler.FetchUsersInRoom)
			rooms.GET("/:id/messages", ChatMessagesHandler.FetchChatMessagesInRoom)
			rooms.POST("", ChatRoomsHandler.CreateRoom)
		}
		likes := api.Group("/messages")
		{
			likes.POST("", ChatMessagesHandler.CreateChatMessage)
			//likes.GET("/:id/likes", ChatLikesHandler.FetchChatLikesInUsers)
			likes.POST("/:id/likes", ChatLikesHandler.CreateChatLike)
			likes.DELETE("/:id/likes", ChatLikesHandler.DeleteChatLike)
		}
	}

	// ---え--------------------------------------------------------
	// WebSocket
	// ------------------------------------------------------------
	// WebSocket 接続
	e.GET("/ws", func(c echo.Context) error {
		WebSocketHandler.HandleWebSocket(c.Response(), c.Request())
		return nil
	})
}
