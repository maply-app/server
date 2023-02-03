package friends

import (
	"github.com/gofiber/fiber/v2"
	"maply/repository/managers/friends"
	"maply/ws"
)

func DeleteFriend(userId, friendId string) error {
	ws.NewEvent(friendId, ws.DeleteFriend, fiber.Map{"id": userId})
	return friends.DeleteFriend(userId, friendId)
}
