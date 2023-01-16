package ws

import (
	"encoding/json"
	"github.com/gofiber/websocket/v2"
	"maply/errors"
)

const (
	SendRequest    = "sendRequest"
	ConfirmRequest = "confirmRequest"
	CancelRequest  = "cancelRequest"
	DeleteFriend   = "deleteFriend"
	FriendsGeo     = "friendsGeo"
)

type Response struct {
	Event string `json:"event"`
	Data  any    `json:"data"`
}

func CreateEvent(eventType string, data any) Response {
	response := Response{}
	response.Event = eventType
	response.Data = data
	return response
}

// NewEvent example –> ws.NewEvent(userID, ws.FriendRequest, user)
func NewEvent(userId, eventType string, msg any) error {
	c := GetClientConnection(userId)
	if c == nil {
		return errors.ObjectDoesNotExists
	}

	event := CreateEvent(eventType, msg)
	bytesMsg, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return c.WriteMessage(websocket.TextMessage, bytesMsg)
}
