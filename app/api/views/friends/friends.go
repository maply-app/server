package friends

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	"maply/core/validators"
	"maply/services/friends"
	"net/http"
)

func DeleteFriend(c *fiber.Ctx) error {
	input := c.Query("userId", "")
	if !validators.UUID(input) || c.Locals("user").(string) == input {
		return core.Send(c, core.Error(core.ValidationError))
	}

	err := friends.DeleteFriend(c.Locals("user").(string), input)
	if err != nil {
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
	return core.Send(c, core.Success(http.StatusOK, nil))
}
