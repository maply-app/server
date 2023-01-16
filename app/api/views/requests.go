package views

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	"maply/api/serializers"
	"maply/errors"
	"maply/services/requests"
	"maply/services/validators"
	"net/http"
)

// GetReceivedRequests ...
func GetReceivedRequests(c *fiber.Ctx) error {
	requests, err := requests.GetReceivedRequests(c.Locals("user").(string))
	if err != nil {
		return core.Send(c, core.Error(core.InternalServerError))
	}
	return core.Send(c, core.Success(http.StatusOK, requests))
}

// GetSentRequests ...
func GetSentRequests(c *fiber.Ctx) error {
	requests, err := requests.GetSentRequests(c.Locals("user").(string))
	if err != nil {
		return core.Send(c, core.Error(core.InternalServerError))
	}
	return core.Send(c, core.Success(http.StatusOK, requests))
}

// SendRequest ...
func SendRequest(c *fiber.Ctx) error {
	input, status := serializers.SendRequestSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	id, err := requests.SendRequest(input)
	switch err {
	case errors.ObjectAlreadyExists:
		return core.Send(c, core.Error(core.ObjectAlreadyExists))
	case nil:
		return core.Send(c, core.Success(http.StatusOK, fiber.Map{"id": id}))
	default:
		return core.Send(c, core.Error(core.ValidationError))
	}
}

// ConfirmRequest ...
func ConfirmRequest(c *fiber.Ctx) error {
	input := c.Query("requestID", "")
	if !validators.UUID(input) {
		return core.Send(c, core.Error(core.ValidationError))
	}

	err := requests.ConfirmRequest(c.Locals("user").(string), input)
	switch err {
	case nil:
		return core.Send(c, core.Success(http.StatusOK, nil))
	default:
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
}

// CancelRequest ...
func CancelRequest(c *fiber.Ctx) error {
	input := c.Query("requestID", "")
	if !validators.UUID(input) {
		return core.Send(c, core.Error(core.ValidationError))
	}

	err := requests.CancelRequest(c.Locals("user").(string), input)
	switch err {
	case nil:
		return core.Send(c, core.Success(http.StatusOK, nil))
	default:
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
}
