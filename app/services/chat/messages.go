package chat

import (
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

func GetMessages(userId, receiverID string, count, offset int) ([]*models.MessageWithoutSender, error) {
	r, err := managers.GetMessages(userId, receiverID, count, offset)
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

func GetChats(userId string, count, offset int) ([]*models.Chat, error) {
	r, err := managers.GetChats(userId, count, offset)
	if err != nil {
		return []*models.Chat{}, err
	}

	var resp []*models.Chat
	for i := range r {
		resp = append(resp, &models.Chat{
			SenderID:   r[i].SenderID,
			ReceiverID: r[i].ReceiverID,
			Text:       r[i].Text,
			CreatedAt:  r[i].CreatedAt,
		})

		var s models.User
		if r[i].SenderID == userId {
			s = r[i].Receiver
		} else {
			s = r[i].Sender
		}

		sender := &models.PublicUserWithoutFriends{}
		deepcopier.Copy(sender).From(s)
		resp[i].Sender = sender
	}
	return resp, nil
}
