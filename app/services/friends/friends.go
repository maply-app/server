package friends

import (
	"github.com/gofiber/fiber/v2"
	"maply/repository/managers"
	"maply/ws"
)

func DeleteFriend(userId, friendId string) error {
	ws.NewEvent(friendId, ws.DeleteFriend, fiber.Map{"id": userId})
	return managers.DeleteFriend(userId, friendId)
}
