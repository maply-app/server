package ws

import (
	"github.com/gofiber/websocket/v2"
)

func baseHandler(c *websocket.Conn) {
	// When the function returns, unregister the client and close the connection
	defer func() {
		unregister <- c
		c.Close()
	}()

	// Register the client
	register <- c

	for {
		_, _, _ = c.ReadMessage()
		//messageType, message, err := c.ReadMessage()
		//if err != nil {
		//	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
		//		fmt.Println("read error:", err)
		//	}
		//	break
		//}
		//
		//if messageType == websocket.TextMessage {
		//	// Broadcast the received message
		//	broadcast <- string(message)
		//} else {
		//	//fmt.Println("websocket message received of type", messageType)
		//}
	}
}
