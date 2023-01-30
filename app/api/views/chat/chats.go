package chat

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	chatSerializers "maply/api/serializers/chat"
	"maply/services/chat"
	"net/http"
)

func GetChats(c *fiber.Ctx) error {
	input, status := chatSerializers.GetChatsSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}
	m, err := chat.GetChats(
		c.Locals("user").(string),
		input.Count,
		input.Offset,
	)
	if err != nil {
		return core.Send(c, core.Error(core.InternalServerError))
	}
	return core.Send(c, core.Success(http.StatusOK, m))
}
