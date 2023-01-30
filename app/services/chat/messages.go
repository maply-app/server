package chat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ulule/deepcopier"
	"maply/errors"
	"maply/models"
	"maply/repository/managers"
	"maply/ws"
	"time"
)

func SendMessage(r *models.Message) error {
	if !managers.CheckFriendByID(r.SenderID, r.ReceiverID) {
		return errors.ObjectDoesNotExists
	}

	// Create a message
	r.ID = uuid.New().String()
	r.CreatedAt = time.Now()
	if err := managers.CreateMessage(r); err != nil {
		return err
	}

	resp := &models.MessageWithSender{}
	resp.Sender = &models.PublicUserWithoutFriends{}
	u, _ := managers.GetUser(r.SenderID)
	deepcopier.Copy(resp.Sender).From(u)
	deepcopier.Copy(resp).From(r)

	// Send socket event
	ws.NewEvent(r.ReceiverID, ws.NewMessage, resp)
	return nil
}

func GetMessages(userId, receiverId string, count, offset int) ([]*models.MessageWithoutSender, error) {
	if offset == 0 {
		if err := managers.ReadMessages(receiverId, userId); err != nil {
			return []*models.MessageWithoutSender{}, err
		}

		// Send socket event
		ws.NewEvent(receiverId, ws.Readessages, fiber.Map{"userId": userId})
	}

	r, err := managers.GetMessages(userId, receiverId, count, offset)
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
	if err := managers.ReadMessages(senderId, userId); err != nil {
		return err
	}

	// Send socket event
	return ws.NewEvent(senderId, ws.Readessages, fiber.Map{"userId": userId})
}
