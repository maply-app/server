package chat

import (
	"github.com/ulule/deepcopier"
	"maply/models"
	chatDBManager "maply/repository/managers/chat"
)

func GetChats(userId string, count, offset int) ([]*models.Chat, error) {
	r, err := chatDBManager.GetChats(userId, count, offset)
	if err != nil {
		return []*models.Chat{}, err
	}

	var resp []*models.Chat
	for i := range r {
		var unRead = 0
		if r[i].SenderID != userId {
			u, err := chatDBManager.GetUnreadMessages(r[i].SenderID, r[i].ReceiverID)
			if err != nil {
				return []*models.Chat{}, err
			}
			unRead = u
		}

		resp = append(resp, &models.Chat{
			SenderID:       r[i].SenderID,
			ReceiverID:     r[i].ReceiverID,
			Text:           r[i].Text,
			UnreadMessages: unRead,
			CreatedAt:      r[i].CreatedAt,
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
