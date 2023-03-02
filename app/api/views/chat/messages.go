package chat

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	chatSerializers "maply/api/serializers/chat"
	"maply/core/validators"
	"maply/errors"
	"maply/services/chat"
	"net/http"
)

func Send(c *fiber.Ctx) error {
	input, status := chatSerializers.SendMessageSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	obj, err := chat.SendMessage(input)
	switch err {
	case errors.ObjectDoesNotExists:
		return core.Send(c, core.Error(core.ObjectNotFound))
	case nil:
		return core.Send(c, core.Success(http.StatusOK, obj))
	default:
		return core.Send(c, core.Error(core.ValidationError))
	}
}

func GetMessages(c *fiber.Ctx) error {
	input, status := chatSerializers.GetMessagesSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}
	m, err := chat.GetMessages(
		c.Locals("user").(string),
		input.ReceiverId,
		input.Count,
		input.Offset,
	)
	if err != nil {
		return core.Send(c, core.Error(core.InternalServerError))
	}
	return core.Send(c, core.Success(http.StatusOK, m))
}

func ReadMessages(c *fiber.Ctx) error {
	input := c.Query("senderID", "")
	if !validators.UUID(input) {
		return core.Send(c, core.Error(core.ValidationError))
	}

	if err := chat.ReadMessages(
		c.Locals("user").(string),
		input,
	); err != nil {
		return core.Send(c, core.Error(core.InternalServerError))
	}
	return core.Send(c, core.Success(http.StatusOK, nil))
}
