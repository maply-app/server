package chat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ulule/deepcopier"
	"maply/errors"
	"maply/models"
	chatDBManager "maply/repository/managers/chat"
	friendsDBManager "maply/repository/managers/friends"
	"maply/services/users"
	"maply/ws"
	"time"
)

func SendMessage(r *models.Message) (*models.MessageWithSender, error) {
	if !friendsDBManager.CheckFriendByID(r.SenderID, r.ReceiverID) {
		return nil, errors.ObjectDoesNotExists
	}

	// Create a message
	r.ID = uuid.New().String()
	r.CreatedAt = time.Now()
	if err := chatDBManager.CreateMessage(r); err != nil {
		return nil, err
	}

	resp := &models.MessageWithSender{}
	resp.Sender = &models.PublicUserWithoutFriends{}
	u, _ := users.GetUser(r.SenderID)
	deepcopier.Copy(resp.Sender).From(u)
	deepcopier.Copy(resp).From(r)

	ws.NewEvent(r.ReceiverID, ws.NewMessage, resp)
	return resp, nil
}

func GetMessages(userId, receiverId string, count, offset int) ([]*models.MessageWithoutSender, error) {
	if offset == 0 {
		if err := chatDBManager.ReadMessages(receiverId, userId); err != nil {
			return []*models.MessageWithoutSender{}, err
		}
		ws.NewEvent(receiverId, ws.Readessages, fiber.Map{"userId": userId})
	}

	r, err := chatDBManager.GetMessages(userId, receiverId, count, offset)
	if err != nil {
		return []*models.MessageWithoutSender{}, err
	}

	var resp []*models.MessageWithoutSender
	for i := range r {
		resp = append(resp, &models.MessageWithoutSender{})
		deepcopier.Copy(resp[i]).From(r[i])
	}
	return resp, nil
}

func ReadMessages(userId, senderId string) error {
	if err := chatDBManager.ReadMessages(senderId, userId); err != nil {
		return err
	}
	ws.NewEvent(senderId, ws.Readessages, fiber.Map{"userId": userId})
	return nil
}
