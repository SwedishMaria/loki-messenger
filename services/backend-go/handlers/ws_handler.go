package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	// upgrader upgrades HTTP connections to a WebSocket connection.
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow all connections by default. For production, you should implement
			// a proper origin check.
			return true
		},
	}
	// clients stores all active WebSocket connections.
	clients = make(map[*websocket.Conn]bool)
	// broadcast is a channel that receives messages and broadcasts them to all clients.
	broadcast = make(chan []byte)
)

// HandleConnections handles incoming WebSocket connections.
// It registers new clients and listens for incoming messages.
func HandleConnections(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	// Register the new client.
	clients[ws] = true

	for {
		// Read a message from the WebSocket connection.
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the message to the broadcast channel.
		broadcast <- msg
	}

	return nil
}

// HandleMessages listens for messages on the broadcast channel and sends them
// to all connected clients.
func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel.
		msg := <-broadcast
		// Send the message to every client.
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}