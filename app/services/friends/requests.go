package friends

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ulule/deepcopier"
	"maply/errors"
	"maply/models"
	"maply/repository/managers/friends"
	"maply/repository/managers/users"
	"maply/ws"
)

func SendRequest(r *models.Request) (string, error) {
	f := friends.CheckFriendByID(r.SenderID, r.ReceiverID)
	if f {
		return "", errors.ObjectAlreadyExists
	}

	req, _ := friends.FindRequestBySenderAndReceiver(r.SenderID, r.ReceiverID)
	if req.ID != "" {
		return "", errors.ObjectAlreadyExists
	}

	// Create a request
	r.ID = uuid.New().String()
	requestID, err := friends.CreateRequest(r)
	if err != nil {
		return "", err
	}

	resp := &models.PrivateRequestWithSender{}
	resp.Sender = &models.PublicUserWithoutFriends{}
	u, _ := users.GetUser(r.SenderID)
	deepcopier.Copy(resp.Sender).From(u)
	deepcopier.Copy(resp).From(r)

	err = ws.NewEvent(r.ReceiverID, ws.SendRequest, resp)
	return requestID, err
}

func GetReceivedRequests(userId string) ([]*models.PrivateRequestWithSender, error) {
	r, err := friends.GetRequestsByReceiver(userId)
	if err != nil {
		return []*models.PrivateRequestWithSender{}, err
	}

	var resp []*models.PrivateRequestWithSender
	for i := range r {
		resp = append(resp, &models.PrivateRequestWithSender{})
		deepcopier.Copy(resp[i]).From(r[i])

		sender := &models.PublicUserWithoutFriends{}
		deepcopier.Copy(sender).From(r[i].Sender)
		resp[i].Sender = sender
	}
	return resp, nil
}

func GetSentRequests(userId string) ([]*models.PrivateRequestWithReceiver, error) {
	r, err := friends.GetRequestsBySender(userId)
	if err != nil {
		return []*models.PrivateRequestWithReceiver{}, err
	}

	var resp []*models.PrivateRequestWithReceiver
	for i := range r {
		resp = append(resp, &models.PrivateRequestWithReceiver{})
		deepcopier.Copy(resp[i]).From(r[i])

		receiver := &models.PublicUserWithoutFriends{}
		deepcopier.Copy(receiver).From(r[i].Receiver)
		resp[i].Receiver = receiver
	}
	return resp, nil
}

func ConfirmRequest(userId, requestID string) error {
	r, err := friends.GetRequestByID(requestID)
	if err != nil {
		return err
	}

	if userId != r.ReceiverID {
		return errors.Forbidden
	}

	err = friends.AddFriend(r.ReceiverID, r.SenderID)
	if err != nil {
		return err
	}

	_, err = friends.DeleteRequestByID(requestID)
	if err != nil {
		return err
	}

	resp := &models.PrivateRequestWithReceiver{}
	resp.Receiver = &models.PublicUserWithoutFriends{}
	u, _ := users.GetUser(r.ReceiverID)
	deepcopier.Copy(resp.Receiver).From(u)
	deepcopier.Copy(resp).From(r)

	ws.NewEvent(r.SenderID, ws.ConfirmRequest, resp)
	return nil
}

func CancelRequest(userId, requestID string) error {
	r, err := friends.GetRequestByID(requestID)
	if err != nil {
		return err
	}

	if err := friends.DeleteRequest(userId, requestID); err != nil {
		return err
	}

	ws.NewEvent(r.ReceiverID, ws.CancelRequest, fiber.Map{"id": requestID})
	ws.NewEvent(r.SenderID, ws.CancelRequest, fiber.Map{"id": requestID})
	return nil
}
