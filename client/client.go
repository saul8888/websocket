package client

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

// Define message
type Message struct {
	Message string `json:"message"`
}

var clients = make(map[*websocket.Conn]bool) // connected clients
var channel = make(chan Message)             // channel channel

var upgrader = websocket.Upgrader{}

func ClientRead(c echo.Context) error {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg) //
		if err != nil {
			log.Printf("Close windows: %v", err)
			delete(clients, ws)
			break
		}

		// Send the newly received message to the channel
		channel <- msg
	}

	return nil
}

func ClientWrite() {
	for {
		//the next message from the channel
		msg := <-channel
		// Send every client that is connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error2: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
