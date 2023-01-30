package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"maply/api/middleware"
	authViews "maply/api/views/auth"
	chatViews "maply/api/views/chat"
	friendsViews "maply/api/views/friends"
	usersViews "maply/api/views/users"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// v1
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	// Auth
	auth := v1.Group("/auth")
	auth.Post("/register", authViews.Register)
	auth.Post("/login", authViews.Login)

	// User
	users := v1.Group("/users", middleware.UserIdentity)
	users.Get("/get", usersViews.Get)
	users.Get("/find", usersViews.Find)
	users.Get("/get-by-id", usersViews.GetByID)

	// Settings
	settings := users.Group("/settings")
	settings.Post("/", usersViews.Settings)

	// Settings
	stats := users.Group("/stats")
	stats.Post("/", usersViews.UpdateStats)

	// Friends
	friends := v1.Group("/friends", middleware.UserIdentity)
	friends.Delete("/delete", friendsViews.DeleteFriend)

	// Requests
	requests := friends.Group("/requests")
	requests.Get("/received", friendsViews.GetReceivedRequests)
	requests.Get("/sent", friendsViews.GetSentRequests)
	requests.Post("/send", friendsViews.SendRequest)
	requests.Get("/confirm", friendsViews.ConfirmRequest)
	requests.Get("/cancel", friendsViews.CancelRequest)

	// Chat
	chats := v1.Group("/chats", middleware.UserIdentity)
	chats.Get("/get", chatViews.GetChats)

	// –> Messages
	messages := chats.Group("/messages", middleware.UserIdentity)
	messages.Post("/send", chatViews.Send)
	messages.Get("/get", chatViews.GetMessages)
	messages.Get("/read", chatViews.ReadMessages)
}
