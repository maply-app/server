package friends

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	friendsSerializers "maply/api/serializers/friends"
	"maply/core/validators"
	"maply/errors"
	"maply/services/friends"
	"net/http"
)

func GetReceivedRequests(c *fiber.Ctx) error {
	r, err := friends.GetReceivedRequests(c.Locals("user").(string))
	if err != nil {
		return core.Send(c, core.Error(core.InternalServerError))
	}
	return core.Send(c, core.Success(http.StatusOK, r))
}

func GetSentRequests(c *fiber.Ctx) error {
	r, err := friends.GetSentRequests(c.Locals("user").(string))
	if err != nil {
		return core.Send(c, core.Error(core.InternalServerError))
	}
	return core.Send(c, core.Success(http.StatusOK, r))
}

func SendRequest(c *fiber.Ctx) error {
	input, status := friendsSerializers.SendRequestSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	id, err := friends.SendRequest(input)
	switch err {
	case errors.ObjectAlreadyExists:
		return core.Send(c, core.Error(core.ObjectAlreadyExists))
	case nil:
		return core.Send(c, core.Success(http.StatusOK, fiber.Map{"id": id}))
	default:
		return core.Send(c, core.Error(core.ValidationError))
	}
}

func ConfirmRequest(c *fiber.Ctx) error {
	input := c.Query("requestID", "")
	if !validators.UUID(input) {
		return core.Send(c, core.Error(core.ValidationError))
	}

	err := friends.ConfirmRequest(c.Locals("user").(string), input)
	switch err {
	case nil:
		return core.Send(c, core.Success(http.StatusOK, nil))
	default:
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
}

func CancelRequest(c *fiber.Ctx) error {
	input := c.Query("requestID", "")
	if !validators.UUID(input) {
		return core.Send(c, core.Error(core.ValidationError))
	}

	err := friends.CancelRequest(c.Locals("user").(string), input)
	switch err {
	case nil:
		return core.Send(c, core.Success(http.StatusOK, nil))
	default:
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
}
