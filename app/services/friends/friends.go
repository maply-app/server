package friends

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ulule/deepcopier"
	"maply/models"
	"maply/repository/managers"
	"maply/ws"
)

// GetFriends ...
func GetFriends(id string) ([]*models.PublicUserWithoutFriends, error) {
	u, err := managers.GetFriends(id)
	if err != nil {
		return []*models.PublicUserWithoutFriends{}, err
	}

	// Pretty response
	var resp []*models.PublicUserWithoutFriends
	for i := range u {
		resp = append(resp, &models.PublicUserWithoutFriends{})
		deepcopier.Copy(resp[i]).From(u[i])
	}
	return resp, nil
}

// DeleteFriend ...
func DeleteFriend(userID, friendID string) error {
	// Send socket event
	ws.NewEvent(friendID, ws.DeleteFriend, fiber.Map{"id": userID})

	return managers.DeleteFriend(userID, friendID)
}
