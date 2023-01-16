package middleware

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	"maply/services/utils"
	"strings"
)

// UserIdentity ...
func UserIdentity(c *fiber.Ctx) error {
	header := string(c.Request().Header.Peek("Authorization"))
	if header == "" {
		return core.Send(c, core.Error(core.Unauthorized))
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" || len(headerParts[1]) == 0 {
		return core.Send(c, core.Error(core.Unauthorized))
	}

	user, err := utils.ParseToken(headerParts[1])
	if err != nil {
		return core.Send(c, core.Error(core.Unauthorized))
	}
	c.Locals("user", user)
	return c.Next()
}
