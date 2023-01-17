package ws

import (
	"fmt"
	"github.com/gofiber/websocket/v2"
)

func newHub() {
	for {
		select {

		case connection := <-register:
			// Register the new client to the hub
			userId := connection.Locals("userId").(string)
			clients[userId] = connection

		case message := <-broadcast:
			// Send the message to all clients
			for id := range clients {
				if err := clients[id].WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					fmt.Println("write error:", err)

					unregister <- clients[id]
					clients[id].WriteMessage(websocket.CloseMessage, []byte{})
					clients[id].Close()
				}
			}

		case connection := <-unregister:
			// Remove the client from the hub
			userId := connection.Locals("userId").(string)
			delete(clients, userId)
		}
	}
}
