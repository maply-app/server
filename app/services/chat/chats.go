package chat

import (
	"github.com/ulule/deepcopier"
	"maply/models"
	"maply/repository/managers"
)

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

		sender := &models.PublicUserWithoutFriends{}
		deepcopier.Copy(sender).From(r[i].Sender)
		resp[i].Sender = sender

		receiver := &models.PublicUserWithoutFriends{}
		deepcopier.Copy(receiver).From(r[i].Receiver)
		resp[i].Receiver = receiver
	}
	return resp, nil
}
