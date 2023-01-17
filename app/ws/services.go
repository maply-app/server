package ws

import "github.com/gofiber/websocket/v2"

func GetClientConnection(userId string) *websocket.Conn {
	return clients[userId]
}
