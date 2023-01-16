package ws

import "github.com/gofiber/websocket/v2"

func GetClientConnection(userID string) *websocket.Conn {
	return clients[userID]
}
