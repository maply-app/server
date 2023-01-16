package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// SetupRoutes setup router ws
func SetupRoutes(app *fiber.App) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if !websocket.IsWebSocketUpgrade(c) {
			return fiber.ErrUpgradeRequired
		}
		if UserIdentity(c) != nil {
			return fiber.ErrUnauthorized
		}
		c.Locals("allowed", true)
		return c.Next()
	})

	// Run the new hub
	go newHub()

	// User controller
	app.Get("/ws/users/controller", websocket.New(baseHandler))
}