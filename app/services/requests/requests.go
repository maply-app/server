package requests

import (
	"github.com/google/uuid"
	"github.com/ulule/deepcopier"
	"maply/errors"
	"maply/models"
	"maply/repository/managers"
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

	r.ID = uuid.New().String()
	return managers.CreateRequest(r)
}

// GetRequests ...
func GetRequests(userID string) ([]*models.PrivateRequest, error) {
	r, err := managers.GetRequestsByReceiver(userID)
	if err != nil {
		return []*models.PrivateRequest{}, err
	}

	// Pretty response
	var resp []*models.PrivateRequest
	for i := range r {
		resp = append(resp, &models.PrivateRequest{})
		deepcopier.Copy(resp[i]).From(r[i])

		sender := &models.PublicUserWithoutFriends{}
		deepcopier.Copy(sender).From(r[i].Sender)
		resp[i].Sender = sender
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
	return nil
}

// CancelRequest ...
func CancelRequest(userID, requestID string) error {
	if err := managers.DeleteRequest(userID, requestID); err != nil {
		return err
	}
	return nil
}
