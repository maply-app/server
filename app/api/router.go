package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"maply/api/middleware"
	"maply/api/views"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// -> v1
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	// -> -> Auth
	auth := v1.Group("/auth")
	auth.Post("/register", views.Register)
	auth.Post("/login", views.Login)

	// -> -> User
	users := v1.Group("/users", middleware.UserIdentity)
	users.Get("/get", views.Get)
	users.Get("/find", views.Find)
	users.Get("/get-by-id", views.GetByID)

	// -> -> Settings
	settings := users.Group("/settings")
	settings.Post("/", views.Settings)

	// -> -> Friends
	friends := v1.Group("/friends", middleware.UserIdentity)
	//friends.Get("/get", views.GetFriends)
	friends.Delete("/delete", views.DeleteFriend)

	// -> -> -> Requests
	requests := friends.Group("/requests")
	requests.Get("/received", views.GetReceivedRequests)
	requests.Get("/sent", views.GetSentRequests)
	requests.Post("/send", views.SendRequest)
	requests.Get("/confirm", views.ConfirmRequest)
	requests.Get("/cancel", views.CancelRequest)
}
