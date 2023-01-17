package views

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	"maply/api/serializers"
	"maply/services/settings"
	"net/http"
)

func Settings(c *fiber.Ctx) error {
	input, status := serializers.SettingsSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	err := settings.Settings(c.Locals("user").(string), input)
	if err != nil {
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
	return core.Send(c, core.Success(http.StatusOK, nil))
}
