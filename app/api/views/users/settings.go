package users

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	usersSerializers "maply/api/serializers/users"
	"maply/services/settings"
	"net/http"
)

func Settings(c *fiber.Ctx) error {
	input, status := usersSerializers.SettingsSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	if err := settings.Settings(c.Locals("user").(string), input); err != nil {
		fmt.Println(err)
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
	return core.Send(c, core.Success(http.StatusOK, nil))
}
