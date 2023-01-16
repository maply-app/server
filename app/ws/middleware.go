package ws

import (
	"github.com/gofiber/fiber/v2"
	"maply/errors"
	"maply/services/utils"
)

// UserIdentity ...
func UserIdentity(c *fiber.Ctx) error {
	token := c.Query("Token")
	if token == "" {
		return errors.Unauthorized
	}

	t, err := utils.ParseToken(token)
	if err != nil {
		return errors.Unauthorized
	}

	c.Locals("userID", t)
	return nil
}
