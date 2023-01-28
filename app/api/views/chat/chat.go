package chat

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	"maply/api/serializers"
	"maply/errors"
	"maply/services/chat"
	"net/http"
)

func Send(c *fiber.Ctx) error {
	input, status := serializers.SendMessageSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	err := chat.SendMessage(input)
	fmt.Printf("err –> %s", err)
	switch err {
	case errors.ObjectDoesNotExists:
		return core.Send(c, core.Error(core.ObjectNotFound))
	case nil:
		return core.Send(c, core.Success(http.StatusOK, nil))
	default:
		return core.Send(c, core.Error(core.ValidationError))
	}
}

func GetMessages(c *fiber.Ctx) error {
	input, status := serializers.GetMessagesSerializer(c)
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

func GetChats(c *fiber.Ctx) error {
	input, status := serializers.GetChatsSerializer(c)
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
