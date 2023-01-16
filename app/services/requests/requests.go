package requests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ulule/deepcopier"
	"maply/errors"
	"maply/models"
	"maply/repository/managers"
	"maply/ws"
)

// SendRequest ...
func SendRequest(r *models.Request) (string, error) {
	f := managers.CheckFriendByID(r.SenderID, r.ReceiverID)
	if f {
		return "", errors.ObjectAlreadyExists
	}

	req, _ := managers.FindRequestBySenderAndReceiver(r.SenderID, r.ReceiverID)
	if req.ID != "" {
		return "", errors.ObjectAlreadyExists
	}

	// Create a request
	r.ID = uuid.New().String()
	requestID, err := managers.CreateRequest(r)
	if err != nil {
		return "", err
	}

	// Pretty response
	resp := &models.PrivateRequestWithSender{}
	resp.Sender = &models.PublicUserWithoutFriends{}
	u, _ := managers.GetUser(r.SenderID)
	deepcopier.Copy(resp.Sender).From(u)
	deepcopier.Copy(resp).From(r)

	// Send socket event
	ws.NewEvent(r.ReceiverID, ws.SendRequest, resp)

	return requestID, err
}

// GetReceivedRequests ...
func GetReceivedRequests(userID string) ([]*models.PrivateRequestWithSender, error) {
	r, err := managers.GetRequestsByReceiver(userID)
	if err != nil {
		return []*models.PrivateRequestWithSender{}, err
	}

	// Pretty response
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

// GetSentRequests ...
func GetSentRequests(userID string) ([]*models.PrivateRequestWithReceiver, error) {
	r, err := managers.GetRequestsBySender(userID)
	if err != nil {
		return []*models.PrivateRequestWithReceiver{}, err
	}

	// Pretty response
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

// ConfirmRequest ...
func ConfirmRequest(userID, requestID string) error {
	r, err := managers.GetRequestByID(requestID)
	if err != nil {
		return err
	}

	if userID != r.ReceiverID {
		return errors.Forbidden
	}

	err = managers.AddFriend(r.ReceiverID, r.SenderID)
	if err != nil {
		return err
	}

	_, err = managers.DeleteRequestByID(requestID)
	if err != nil {
		return err
	}

	// Pretty response
	resp := &models.PrivateRequestWithReceiver{}
	resp.Receiver = &models.PublicUserWithoutFriends{}
	u, _ := managers.GetUser(r.ReceiverID)
	deepcopier.Copy(resp.Receiver).From(u)
	deepcopier.Copy(resp).From(r)

	// Send socket event
	ws.NewEvent(r.SenderID, ws.ConfirmRequest, resp)

	return nil
}

// CancelRequest ...
func CancelRequest(userID, requestID string) error {
	r, err := managers.GetRequestByID(requestID)
	if err != nil {
		return err
	}

	if err := managers.DeleteRequest(userID, requestID); err != nil {
		return err
	}

	// Send socket event
	ws.NewEvent(r.ReceiverID, ws.CancelRequest, fiber.Map{"id": requestID})
	ws.NewEvent(r.SenderID, ws.CancelRequest, fiber.Map{"id": requestID})

	return nil
}
