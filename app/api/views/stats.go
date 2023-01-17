package views

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	"maply/api/serializers"
	"maply/services/stats"
	"net/http"
)

func UpdateStats(c *fiber.Ctx) error {
	input, status := serializers.StatsSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	userId := c.Locals("user").(string)
	err := stats.UpdateStats(userId, input)
	if err != nil {
		return core.Send(c, core.Error(core.InternalServerError))
	}
	
	stats.GetStats(userId)
	return core.Send(c, core.Success(http.StatusOK, nil))
}
