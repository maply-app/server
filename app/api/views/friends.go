package views

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	"maply/services/friends"
	"maply/services/validators"
	"net/http"
)

// GetFriends ...
func GetFriends(c *fiber.Ctx) error {
	input := c.Query("userID", "")
	if !validators.UUID(input) {
		return core.Send(c, core.Error(core.ValidationError))
	}

	friends, err := friends.GetFriends(input)
	if err != nil {
		return core.Send(c, core.Error(core.InternalServerError))
	}
	return core.Send(c, core.Success(http.StatusOK, friends))
}

// DeleteFriend ...
func DeleteFriend(c *fiber.Ctx) error {
	input := c.Query("userID", "")
	if !validators.UUID(input) || c.Locals("user").(string) == input {
		return core.Send(c, core.Error(core.ValidationError))
	}

	err := friends.DeleteFriend(c.Locals("user").(string), input)
	if err != nil {
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
	return core.Send(c, core.Success(http.StatusOK, nil))
}
