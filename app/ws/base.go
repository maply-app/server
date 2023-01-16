package ws

import "github.com/gofiber/websocket/v2"

var clients = make(map[string]*websocket.Conn)
var register = make(chan *websocket.Conn)
var broadcast = make(chan string)
var unregister = make(chan *websocket.Conn)
